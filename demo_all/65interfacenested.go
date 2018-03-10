package main

import "fmt"

type USB interface {
	getName() string //定义一个方法，参数为空，返回string类型
	/*抽取到 Standard 接口 
	connect()
	disconnect() */
	Standard
}

type Standard interface {
	connect()
	disconnect()
}


type PCWire struct {
	name string
}

func main() {
	var usb USB
	usb = PCWire{"PC 数据线"} //usb的运行时类型是 struct PCWire
	str := usb.getName()
	fmt.Println(str) //PC 数据线

	//一般的用法，不用定义接口实例，直接指定实现接口的类型实例
	usb2 := PCWire{"劣质 PC 数据线"}
	usb2.connect()
	usb2.disconnect()

	usb3 := PCWire{"高级 PC 数据线"}
	(*PCWire).connect(&usb3)
	(*PCWire).disconnect(&usb3)


}

//实现了usb意味着三个方法都要实现

//receiver 可以指定对象或者指针
func (wire PCWire) getName() string {
	return wire.name
}

func (wire PCWire) connect() {
	fmt.Println(wire.name + "已经连接 USB")
}

func (wire PCWire) disconnect() {
	fmt.Println(wire.name + "已经断开USB")
}
