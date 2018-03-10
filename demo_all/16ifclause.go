package main

import "fmt"

func main()  {
	a := 1
	
	//不带括号
	if a > 0 {
		fmt.Println("a 可以继续参加运算.");
	}

	//支持一个初始化表达式, 注意是分号
	if b:=1; a == b {
		fmt.Println("a 和 b 相等.");
	}

	//b是局部变量，出了上述作用于范围就嗝屁了
	//fmt.Println(b); //编译报错: undefined: b
}