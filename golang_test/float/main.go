package main

import "fmt"

func main() {
	var a float64 = 0.1234567894
	a += 0.0 / 3600.0
	fmt.Println(a)
}
