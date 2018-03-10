package main

import "fmt"

func main() {

	/* 数组指针 */
	a := [10]int{2:1, 3:1}
	//声明一个数组指针(写完数组类型，前面补一个 *号即可)
	var p *[10]int
	//p = a //编译报错:cannot use a (type [10]int) as type *[10]int in assignment
	p = &a //数组是值类型, 可以把它认为是常量

	fmt.Println(p)
	fmt.Println(*p)
}