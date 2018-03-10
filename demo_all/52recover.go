package main

import "fmt"

func main() {
	//碰到 panic 后面基本就终止了
	defer fmt.Println("我只处理main函数内的异常")
	A()
	B()
	C()

}

func A() {
	fmt.Println("func A")
}

func B() {
	fmt.Println("func B")

	//引发问题前做好异常处理
	defer func() {
		if flag := recover(); flag != nil {
			//表明确实引发了异常
			fmt.Println("函数B能引发了异常，我来恢复")
		}
	}()

	//下面引发一个panic
	panic("发生严重错误了") //制定了recover就不会panic了
}

func C() {
	fmt.Println("func C")
}