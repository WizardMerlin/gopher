package main

import "fmt"


type XX int

func main() {
	var x XX
	//x.printInt()

	//通过 method express调用
	(*XX).printInt(&x) //传入实例的指针
}

//给 XX 指定方法

func (p *XX)printInt() {
	*p = *p + 1 //因为 X 就是 int
	fmt.Println(*p)
}