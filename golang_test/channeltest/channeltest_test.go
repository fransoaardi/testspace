package main

import (
	"testing"
	"fmt"
)

// TestChannelOne tests getting value from closed channel.
// It is possible to get value from closed channel.
// It returns zero and false value when nothing is left in closed channel.
// It panics when putting value into closed channel.
func TestChannelOne (t *testing.T) {

	test := make (chan int ,100)

	for i:=0; i<10; i++{
		test <- i
	}
	// test channel is now closed
	// using range without closing channel panics
	close(test)

	for i:=0; i<15; i++{
		val, ok := <-test
		fmt.Println(val, ok)
	}

}

// TestChannelTwo tests using for range with closed channel.
// It iterates until it gets every value from closed channel.
func TestChannelTwo (t *testing.T){

	test := make (chan int ,100)

	for i:=0; i<10; i++{
		test <- i
	}
	// test channel is now closed
	close(test)

	for val := range test{
		fmt.Println(val)
	}
	val,ok := <-test
	fmt.Println(val,ok)
	val,ok = <-test
	fmt.Println(val,ok)

	// It panics
	// test <- 11
}