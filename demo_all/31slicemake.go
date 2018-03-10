package main

import "fmt"

func main() {
	//使用 make 方法
	slice := make([]int, 10, 15) //指定长度为10，容量为15
	fmt.Println(slice) //打印整个 slice
	fmt.Println(len(slice), cap(slice)) //10, 15
}