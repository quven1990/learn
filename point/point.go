package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	a = true
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

}
