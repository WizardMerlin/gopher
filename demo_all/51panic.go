package main

import "fmt"

func main() {
	//碰到 panic 后面基本就终止了
	defer fmt.Println("先声明")
	A()
	B()
	C()
	defer fmt.Println("后声明") //后声明的没有执行

}

func A() {
	fmt.Println("func A")
}

func B() {
	fmt.Println("func B")
	//下面引发一个panic
	panic("发生严重错误了")
}

func C() {
	fmt.Println("func C")
}