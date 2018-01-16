package main

import (
	"fmt"
	"time"
	"sync"
)
var(
	wg = sync.WaitGroup{}
)
func main(){
	wg.Add(1)
	ch := make(chan int,2)	// make buffer small
	go thrower(ch)
	go catcher(ch)
	wg.Wait()
}

/*
	When buffer reaches its size,
	thrower sleeps and wait until buffer releases

 */

func thrower(ch chan<- int){
	for i:=0; i<10; i++{
		ch <- i
		fmt.Printf("Throwed >>  %d\n",i)
		time.Sleep(100*time.Millisecond)	// throws fast , catches slow
	}
}

func catcher(ch <-chan int){
	defer wg.Done()
	for j:=0; j<10; j++{
		i := <-ch
		fmt.Printf("\t\t\t%d << Catched \n",i)
		time.Sleep(2000*time.Millisecond)
	}
}