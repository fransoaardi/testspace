package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var sm sync.Map

	sm.Store("key1", "value1")
	sm.Store("key2", "value2")
	sm.Store("key3", "value3")
	sm.Store("key4", "value4")

	for i := 0; i <= 4; i++ {
		val, ok := sm.Load("key" + strconv.Itoa(i))
		if !ok {
			fmt.Println("no value inside")
		} else {
			fmt.Println("key"+strconv.Itoa(i)+" : ", val)
		}
	}

}
