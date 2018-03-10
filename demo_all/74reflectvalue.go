package main

import "fmt"
import "reflect"

func main() {
	x := 99
	fmt.Println(x)

	//先获得值得反射对象
	val := reflect.ValueOf(&x) //必须传递指针，因为下面要修改原值
	val.Elem().SetInt(100) 
	//btw: Elem returns the value that the interface v contains

	fmt.Println(x)

}