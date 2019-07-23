package main

import "fmt"

type ESlice struct{
	num int
	errors []error
}

func New() *ESlice{
	return &ESlice{
		num: 1,
	}
}

func main(){
	var errorSlice []error
	for i:=0; i<5; i++{
		errorSlice = append(errorSlice, nil)
	}
	fmt.Println(len(errorSlice))
	fmt.Println(errorSlice)

	es := New()
	for i:=0; i<0; i++{
		es.errors = append(es.errors, nil)
	}
	fmt.Println(len(es.errors))
	fmt.Println(es.errors)
}
