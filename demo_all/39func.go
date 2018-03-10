package main

import "fmt"

func main() {
	a, b := f(1)
	fmt.Println(a, b)

	c, d := f1()
	fmt.Println(c,d)
}

func f(param1 int)(result1, result2 int) {
	result1, result2 = param1, 0
	return //因为已经给返回值命名了，系统知道要返回的是result1, reuslt2
}

func f1()(int, int) {
	//这里返回值没有命名, 所以下面要指定
	result1, result2 := 1, 0
	return result1, result2
}