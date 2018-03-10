package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("main函数: x地址 = %p\n", &x) //0xc420012090
	codeBlock := closure(x)
	codeBlock(3.14)
}

//探究一下 闭包捕获的变量是不是原始变量? 看来默认值拷贝。

func closure(x int) func(float32)float32 {
	fmt.Printf("闭包内部: x地址 = %p\n", &x) //0xc420012098
	return func(y float32) float32 {
		fmt.Printf("Code Block内部: x地址 = %p\n", &x) //0xc420012098
		return y
	}
}