package main

import "fmt"

type person struct {
	name string
	age int
}

type person1 struct {
	name string
	age int
}

func main() {
	pp0 := person {
		name: "bob",
		age:20,
	}
	pp1 := person {
		name: "bob",
		age:20,
	}
	fmt.Println("pp0 == pp1, ", pp0 == pp1) //true

	pp1.age += 1
	fmt.Println("pp0 == pp1, ", pp0 == pp1) //false

	//bb0 := person1 {"bob", 20}
	//fmt.Println("bb0 == pp0, ", bb0 == pp0) //编译报错: 
	//mismatched types person1 and person

}