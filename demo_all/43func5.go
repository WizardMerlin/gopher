package main

import "fmt"

func main() {
	//匿名函数的调用
	a := func () {
		fmt.Println("我是匿名函数")
	}
	a()
}