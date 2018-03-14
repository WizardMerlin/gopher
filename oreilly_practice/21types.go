package main

import ( 
	"fmt"
)

func whatisThis(i interface{}) {
	//fmt.Printf("%T\n", i)
	switch i.(type) {
	case string:
		fmt.Println("a string")
	case int:
		fmt.Println("a integer", i.(string))
	default:
		fmt.Printf("Don't know\n")
	}
}

func main() {
	whatisThis(10)
}