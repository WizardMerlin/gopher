package main

import "fmt"
//匿名字段
type person struct {
	string
	int
}

func main() {
	//way1 按顺序初始化
	pp := &person {"mike", 20}
	//pp := &person {20, "mike"}//编译报错
	fmt.Println(*pp)

	//way2 ???
	
}