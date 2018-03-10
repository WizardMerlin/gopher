package main

import "fmt"

func main() {
	//定义的同时初始化

	pp := &struct {
		name string
		age int
	}{
		name: "bob",
		age: 12,
	}
	fmt.Println(*pp)
}