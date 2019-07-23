package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CaseNewBytesBuffer(bf *bytes.Buffer){
	fmt.Println("Case : [NewBytesBuffer]")
	req, err := http.NewRequest("POST", "http://localhost:8080/url-here", bf)
	if err != nil {
		fmt.Println(err)
	}

	r1, err := req.GetBody()
	if err != nil {
		fmt.Println(err)
	}

	read, err := ioutil.ReadAll(r1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read GetBody", string(read))

	read, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read Body", string(read))

	r2, err := req.GetBody()
	if err != nil {
		fmt.Println(err)
	}

	read, err = ioutil.ReadAll(r2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read GetBody2", string(read))

	read, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read Body2", string(read))

}

func CaseNewBytesReader(br *bytes.Reader){
	fmt.Println("Case : [NewBytesReader]")
	req, err := http.NewRequest("POST", "http://localhost:8080/url-here", br)
	if err != nil {
		fmt.Println(err)
	}

	r1, err := req.GetBody()
	if err != nil {
		fmt.Println(err)
	}

	read, err := ioutil.ReadAll(r1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read GetBody", string(read))

	read, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read Body", string(read))

	r2, err := req.GetBody()
	if err != nil {
		fmt.Println(err)
	}

	read, err = ioutil.ReadAll(r2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read GetBody2", string(read))

	read, err = ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Read Body2", string(read))

}

func main() {
	var jsonString= `{"name":"john"}`
	testByte := []byte(jsonString)

	nbf := bytes.NewBuffer(testByte)

	fmt.Println("\nbuffer before : ",nbf)
	CaseNewBytesBuffer(nbf)
	fmt.Println("buffer after : ",nbf)


	nbr := bytes.NewReader(testByte)
	fmt.Println("\nreader before : ",nbr)
	CaseNewBytesReader(nbr)
	fmt.Println("reader after : ",nbr)
}