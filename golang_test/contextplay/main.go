package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
func do() error {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err!=nil {
		return err
	}

	ctx, cancel := context.WithTimeout(req.Context(), 10000 * time.Millisecond )
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err!=nil{
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return err
	}
	fmt.Println(string(body))

	return nil
}

func main(){
	do()
}

