package main

import (
	"io"
	"net/http"
	"fmt"
)

func fakeSlowHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
	//time.Sleep(5000 * time.Millisecond)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Header.Get("X-Forwarded-For"))
}
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/test",fakeSlowHandler)

	server := http.Server{
		Addr : "localhost:9090",
		Handler: mux,
	}
	server.ListenAndServe()
}
