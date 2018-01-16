package main

import (
	"fmt"
	"net/http"
	"time"
	_ "net/http/pprof"
	"sync/atomic"
	"sync"
)

var hit = [30000]int{}

var jobs = make(chan int32)
var count int32
var processed int32
var mut sync.Mutex

func pool (){
	for i:=1; i<=3; i++{	// make three workers
		go worker(i)
	}
}

func worker(idx int) {

	for {
		select{
		case p := <-jobs:
			atomic.AddInt32(&processed,1)
			fmt.Println("Processing job : ", idx, p)
			hit[p]++
			time.Sleep(5000*time.Millisecond)
			fmt.Println("Done job : ", idx, p)
		}
	}
}
func worker2(a int32){
	atomic.AddInt32(&processed,1)

	fmt.Println("Processing job : ", a)
	hit[a]++
	time.Sleep(5000*time.Millisecond)
	fmt.Println("Done job : ", a)
}

func pool2(){
	for{
		go worker2(<-jobs)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count,1)
	jobs <- count

	fmt.Println(time.Now(),"@@@@@@@@@@@@@@@@@@@@@@@@@@@",count)

	fmt.Fprintln(w, "hello world",count, processed)
}
func resultHandler(w http.ResponseWriter, r *http.Request) {
	var more = []int{}
	var empty = []int{}
	for i:=1; i<=int(count); i++{
		switch {
			case hit[i]==0:
				empty = append(empty,i)
			case hit[i]>1:
				more = append(more,i)
		}
	}
	fmt.Fprintf(w,"empty(%v) : %v \n more(%v) : %v",cap(empty), empty,cap(more),more)
}

func main() {
	go pool()
	//go pool2()

	http.HandleFunc("/async", handler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe(":9090", nil)
}
