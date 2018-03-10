package main

import "fmt"

type empty struct {}

type person struct {
	name string
	age int
}

func main() {
	ee := empty{}
	fmt.Println(ee)

	pp := person {}
	fmt.Println(pp)
	pp.name = "bob"
	pp.age = 20
	fmt.Println(pp)

	pp1 := person{"bob",22, //直接字面值指定
	} //字面值直接初始化(括号必须换行)
	fmt.Println(pp1)

	pp2 := person{ //指定索引
		name:"mike",
		age:23, //直接字面量初始化
	}
	fmt.Println(pp2)
}