package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []interface{}{true, 1, "yes"}
	e := s[0]
	fmt.Println(e, reflect.TypeOf(e)) // true bool

	//but when used this way:
/* 	if e {
		fmt.Println("This value is true")
	} //complains : non-bool e (type interface {}) used as if condition */
	
	//way one: type assertion
	if e.(bool) {
		fmt.Println("This value is true")
	} //complains : non-bool e (type interface {}) used as if condition

	//way two: using == to get the runtime type
	if e == true {
		fmt.Println("This value is true too");
	}
}