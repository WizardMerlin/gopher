package main

import "fmt"

func main()  {

	var a rune = 10
	//var b int

	//b = a //编译报错 : cannot use a (type rune) as type int in assignment
	b := int(a) //强制类型转换

	fmt.Println(b)
}