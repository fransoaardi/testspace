package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	test := `{"aa":"bb"}`
	var class interface{}
	err := json.Unmarshal([]byte(test), &class)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%#v",class)
}
