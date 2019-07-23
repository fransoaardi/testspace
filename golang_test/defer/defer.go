package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer2()

}

func defer1() {
	//defer func(){
	//	if err:=recover(); err!=nil{
	//		fmt.Println(err)
	//	}
	//}()
	panic("panic : defer1()")
}

func defer2() {
	//defer func(){
	//	if err:=recover(); err!=nil{
	//		fmt.Println(err)
	//	}
	//}()
	defer1()

}
