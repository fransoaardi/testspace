package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan struct{})
	quit := false

	ticker := time.NewTicker(1*time.Second)

	go quitInThreeSeconds(ch)

EXIT:
	for {
		select{
		case <-ch:
			quit = true
			fmt.Println("break select, for loop again")
			break
		case <-ticker.C:
			if quit{
				fmt.Println("break for loop. In golang, go to after for loop")
				break EXIT
			}
		}
	}
	fmt.Println("break", quit)
}

func quitInThreeSeconds(ch chan struct{}){
	time.Sleep(3*time.Second)
	ch <- struct{}{}
	close(ch)
	return
}
