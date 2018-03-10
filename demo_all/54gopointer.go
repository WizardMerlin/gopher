package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	//默认的传参数，都是值传递；不说

	//way1, 先构建一个 person 对象，然后取地址
	p1 := person {
		name: "bob",
		age: 20, //不要忘记逗号
	}
	setPerson(&p1)
	fmt.Println("outer: ", p1)


	//way2 -- 直接拿到struct的指针
	p2 := &person {
		name: "mike",
		age:30,
	}
	setPerson(p2)
	fmt.Println("outer: ", *p2)//这里p2是指针

	//指针也用 . 操作
	p2.age = 32
	fmt.Println("outer: ", *p2)
}

func setPerson(p *person){
	p.age += 1
	fmt.Println("inner: ", *p) 
}