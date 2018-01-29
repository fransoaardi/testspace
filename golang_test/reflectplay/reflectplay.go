package main

import (
	"fmt"
	"reflect"
)

func main(){

	t:= reflect.TypeOf(3)
	v:= reflect.ValueOf(3)

	fmt.Println(t, v)

	t = reflect.TypeOf("ABCDE")	// reflect.Type
	fmt.Printf("%T \n","ABCDE")		// %T prints reflect.TypeOf()

	v = reflect.ValueOf("ABCDE")	// reflect.Value

	fmt.Printf("%v \n","ABCDE")		// %v prints reflect.ValueOf()
	fmt.Println(v.Interface())		// reflect.Value.Interface() returns origin( interface{} )

	fmt.Println(t,v)

}
