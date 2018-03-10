package main

import "fmt"

type USB interface {
	printName()
}

type PCWire struct {
	name string
}

func main() {
	//var wire USB
	//接口 = 对象, 此处就发生了拷贝，即 wire 其实 struct PCWire 实例的副本
	wire := PCWire{"电脑数据线"} 
	var usb USB //接口内部保存着对于实例的引用，所以拿到该部分引用(原始对象)
	usb = USB(wire) //抽象接口可以兼容，拿到接口所存储的实例的指针

	wire.printName() //副本打印 "电脑数据线"
	usb.printName() //接口拿到的打印 "电脑数据线" 

	//修改 wire 这个副本的内容
	wire.name = "数据线(改)"
	//当然通过 usb.name 修改肯定不行啦，接口没有这个字段

	wire.printName() //副本的改变了: 数据线(改)
	usb.printName()  //接口内存储的指针则没有变: 电脑数据线
}

func (wire PCWire) printName() {
	fmt.Println(wire.name)
}

/*
电脑数据线
电脑数据线
数据线(改)
电脑数据线
*/