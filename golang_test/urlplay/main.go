package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	req, err  := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil{
		fmt.Println(err)
	}
	val := req.URL.Query()
	val.Add("q", "아이유" )
	req.URL.RawQuery = val.Encode()
	fmt.Println(req)
	req.FormValue("da")
	fmt.Println(req)
	fmt.Println(req.FormValue("q"))
	fmt.Println(req.FormValue("da"))

	uv := url.Values{
		"q": {"아이유"},
		"q1": {"john"},
	}

	pReq, err := http.NewRequest("POST", "http://www.google.com",  ioutil.NopCloser(bytes.NewBufferString(uv.Encode())))
	if err != nil{
		fmt.Println(err)
	}
	pReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(pReq)
	fmt.Println(pReq.Body)

	// Note: when the request is POST and application/x-www-form-urlencoded, FormValue consumes all the body and makes values with map
	fmt.Println(pReq.FormValue("q"))
	fmt.Println(pReq.Body)
}
