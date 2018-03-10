package main

import "fmt"

type USB interface {
	getName() string //定义一个方法，参数为空，返回string类型
	connect()
	disconnect()
}

//定义一个结构，PC数据线，实现了USB接口
type PCWire struct {
	name string
}

func main() {
	var usb USB
	usb = PCWire{"PC 数据线"} //usb的运行时类型是 struct PCWire
	//usb.name undefined (type USB has no field or method name)
	//usb.name = "PC 数据线"//应该在上一步赋值的时候指定
	str := usb.getName()
	fmt.Println(str) //PC 数据线

	//一般的用法，不用定义接口实例，直接指定实现接口的类型实例
	usb2 := PCWire{"劣质 PC 数据线"}
	usb2.connect()
	usb2.disconnect()

	usb3 := PCWire{"高级 PC 数据线"}
	(*PCWire).connect(&usb3)
	(*PCWire).disconnect(&usb3)

	fmt.Println("-----")
	ensureInterface(usb)

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

//来一个形参方法，接收 USB 实例 (面向接口编程)
func ensureInterface(usb USB) {
	fmt.Println("正常调用该方法，说明传入的实参是 usb 的实例")
	fmt.Println(usb.getName()) //直接 usb.name 不行，usb没有这个字段
	//编译报错:cannot convert usb (type USB) to type PCWire: need type assertion
	//fmt.Println(PCWire(usb).name) //强制类型转换？
	if pcWire,ok := usb.(PCWire); ok == true {
		//如果 usb 确实是 PCWire 的实例
		fmt.Println("Disconnected: ", pcWire.name)
	}
}