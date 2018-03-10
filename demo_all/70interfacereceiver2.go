package main

//证明一下 receiver 不能自动转换

import "fmt"

type print interface {
	printPerson1()
	printPerson2()
}

type person struct {
	name string
	age int
}

func main() {
	//这里pp0,pp1分别是接口类型副本，指针
	pp0 := person{"bob",20} 
	pp1 := &person{"bob", 20}

	pp0.printPerson1()
	pp0.printPerson2()


	pp1.printPerson1()
	pp1.printPerson2()
}
//该接口方法，实例和指针均可调用，自动转换ok
func (p person)printPerson1() {
	fmt.Println(p.name, p.age)
}

//该接口方法，只能指针调用，实例调用失败----事实证明可以！！！
func (p *person)printPerson2() {
	fmt.Println(p.name, p.age)
}