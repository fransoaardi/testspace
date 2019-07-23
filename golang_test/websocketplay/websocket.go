package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{"http://daum.net","http://daum.net","http://daum.net","http://daum.net"}

	for _, val := range urls{
		req, err := http.NewRequest("GET", val, nil)
		if err!=nil{
			fmt.Println(err)
		}
	}

	for {
		select{
		}
	}
}

func getResp(req *http.Request) *Response{
	resp, _ := http.DefaultClient.Do(req)

}