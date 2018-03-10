package main

import "fmt"

func main() {

	// A 函数调用了 AA 函数
	A()
}

func A() {
	a := AA
	a()  //相当于函数调用 AA()
}

func AA() {
	fmt.Println("调用AA")
}