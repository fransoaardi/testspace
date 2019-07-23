package main

import (
	"fmt"
	"strings"
)

func main() {
	inString := "key\tval\tone\ttwo\tthree"
	elem := strings.Split(inString, "\t")
	fmt.Println(elem[0])
	fmt.Println(elem[1:])
	fmt.Println(strings.Join(elem[1:], ""))

	fmt.Println(elem)
	elem = elem[0:0]
	fmt.Println(elem)
}