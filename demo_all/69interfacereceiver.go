package main

//证明一下 receiver 不能自动转换

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	pp0 := person{"bob",20}
	pp1 := &person{"bob", 20}

	//访问自动，自动转换 ok
	pp0.age++
	pp1.age++

	//访问方法自动转换 ok ( 调用 printPerson2() 把 printPerson1() 注释掉 )
/* 	pp0.printPerson2()
	pp1.printPerson2() */

	//调用 printPerson1() 把 printPerson2() 注释掉
	pp0.printPerson1()
	pp1.printPerson1()
}

func (p person)printPerson1() {
	fmt.Println(p.name, p.age)
}

/* func (p *person)printPerson2() {
	fmt.Println(p.name, p.age)
} */