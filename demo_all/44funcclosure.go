package main

import "fmt"

func main() {
	//闭包返回一段代码块
	codeBlock := closure(1) //返回类型 codeBlock 的类型是 func(int) int
	result1 := codeBlock(0)
	result2 := codeBlock(2)
	fmt.Println("result1 = ", result1)
	fmt.Println("result2 = ", result2)
}

func closure(x int) func(int) int {
	//closre(x)闭包返回 func(int) int 代码块
	return func(y int) int {
		fmt.Println("x = ", x)
		fmt.Println("y = ", y)
		return x+y
	}
}