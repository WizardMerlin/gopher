package main

import "fmt"

func main() {
	var a [2]int
	//var b [1]int //长度不同，不能相互赋值
	var b [2]int //长度相同则可以

	a = b //长度不同，编译报错: cannot use b (type [1]int) as type [2]int in assignment
	fmt.Println(a);

	//如果不默认赋值，那么系统自动补零值
	//var aa int[3]
	aa := [3]int{1}
	fmt.Println(aa)

	//也可以按照索引指定(索引后面跟着冒号)
	bb := [5]int{3:1} //第4个元素为1，其他为0
	fmt.Println(bb)

	//逐个指定
	cc := [3]int{1:2, 2:3}
	fmt.Println(cc)


	//...指定长度 (根据后面初始化的元素，确定长度)
	// 如果指定索引，那么也会根据索引来推断长度
	words := [...]string{"hello", "world"}
	for _, word := range words {
		fmt.Println(word)
	}
}