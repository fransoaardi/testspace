// curl 'https://localhost:9000/test' --tlsv1.3 --http2 -kIv
// needs -k flag because certificate and key file is self-made which is insecure.
// server using tls configuration in golang transparently serves http2

package main

import (
	"fmt"
	"net/http"
	"os"
)

func init() {
	// to enable tlsv1.3 (go 1.12)
	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",tls13=1")
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/test", TestHandler)

	server := http.Server{
		Addr: "localhost:9000",
		Handler: mux,
	}

	if err := server.ListenAndServeTLS("keys/server.crt", "keys/server.key"); err != nil{
		fmt.Println(err)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Test called")
	fmt.Println(r)
	w.Write([]byte("hi hello"))
}
