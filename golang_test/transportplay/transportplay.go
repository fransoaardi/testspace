package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

var transport http.Transport
var atom int32

func init(){
	transport = http.Transport {
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func newClient(timeout time.Duration) *http.Client{
	return &http.Client{
		Transport:&transport,
		Timeout: timeout,
	}
}

func startWebserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("Hello, %v\n", r.URL.Query().Get("cnt"))
	})
	go http.ListenAndServe(":8080", nil)
}

func startLoadTest() {
	client := newClient(1100*time.Millisecond)
	//client := &http.Client{Timeout:1000*time.Millisecond}
	for {
		now := time.Now()
		cnt := atom
		//log.Printf("Started GET request #%v %v", cnt, now)
		resp, err := client.Get("http://localhost:8080?cnt="+fmt.Sprint(cnt))
		if err != nil {
			panic(fmt.Sprintf("Got error: %v , atom : %v %v %v %v", err, cnt, now, time.Now(), time.Since(now)))
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Printf("Finished GET request #%v %v %v", cnt, time.Now(), time.Since(now))
		atomic.AddInt32(&atom, 1)
	}

}

func main() {

	startWebserver()
	for i := 0; i < 10; i++ {
		go startLoadTest()
	}

	time.Sleep(time.Second * 2400)
}
