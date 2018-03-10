package main

import "fmt"


func main()  {
	var b int
	b = 1 //ok
	b = 1.1 //编译报错: constant 1.1 truncated to integer
	fmt.Println(b)
}