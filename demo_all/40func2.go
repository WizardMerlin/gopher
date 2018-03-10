package main

import "fmt"

func main() {
	f(1,2,3,4)

	//变参实际上是值拷贝
	a, b, c := 1, 2, 3
	f1(a, b, c)
	fmt.Println("外部: ", a, b, c)
}

//不定长变参函数 (必须作为最后一个参数, 编译器的要求)
func f(a ...int) { //没有返回值，可以直接不写
	fmt.Println(a)
}

func f1(a ...int) {
	//如果传递的是值，那么内部修改并不会影响外部的值
	a[0] = 3
	a[1] = 2
	a[2] = 1
	fmt.Println("内部: ", a)
}

