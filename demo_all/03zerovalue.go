package main

import "fmt"

func main() {
	//验证一下各种类型的零值
	var cnt int
	fmt.Println(cnt) //0

	var bb []byte
	fmt.Println(bb) //[]


	var intArr [3]int  //不指定大小就是 slice
	fmt.Println(intArr) //[0,0,0]

	var flag bool
	fmt.Println(flag) //false
}