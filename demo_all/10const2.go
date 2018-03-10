package main

import "fmt"

var str = "xxxx" //把常量赋值给变量str

func main()  {

	//length := len(str) //ok
	const length = len(str) //编译报错 : const initializer len(str) is not a constant
	fmt.Println(length)
}