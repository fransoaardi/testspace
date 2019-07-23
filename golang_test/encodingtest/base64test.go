package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	input := "abcde"
	fmt.Println(input)

	beforeEncode := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println(beforeEncode)

	afterEncode, _ := base64.StdEncoding.DecodeString(beforeEncode)
	fmt.Println(string(afterEncode))
}
