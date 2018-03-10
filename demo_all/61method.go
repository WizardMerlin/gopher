package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	pp := person{"bob", 10}
	fmt.Println("old: pp", pp)

	pp.printPerson1()
	fmt.Println("outer print1: pp", pp) //内部改变，但外部没有变

	pp.printPerons2()
	fmt.Println("outer print2: pp", pp)
}
//指定 receiver
func (p person) printPerson1() { //默认按值拷贝
	p.age ++
	fmt.Println("inner print1", p.name, p.age)
}

func (p *person) printPerons2() { //指定传递 person 对象的指针
	p.age ++
	fmt.Println("inner print2", p.name, p.age)
}

/* 输出结果
old: pp {bob 10}
inner print1 bob 11
outer print1: pp {bob 10}
inner print2 bob 11
outer print2: pp {bob 11}
*/