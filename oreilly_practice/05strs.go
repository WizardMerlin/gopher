package main

import "fmt"

func main() {

	str := "hello,world,hellowold"

	substr1 := str[1:]
	substr2 := str[1:2]

	fmt.Println(substr1)
	fmt.Println(substr2)

	for i, r := range str {
		fmt.Println(i, " : ", r)
	}
 }