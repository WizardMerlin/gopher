package main

import "fmt"

func main() {
	defer fmt.Println("执行 1")
	defer fmt.Println("执行 2")
	fmt.Println("main函数执行")
	A()
}

func A() {
	fmt.Println("执行 3")
	fmt.Println("子函数执行")
}