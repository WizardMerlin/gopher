package main

import "fmt"

type A struct {
	name string
}

type B struct {
	name string
}

type C1 struct {
	A
	B
}

type C2 struct {
	A
	B
	name string
}

func main() {
	cc1 := C1{} //ok
	fmt.Println(cc1)

	//编译报错
	//fmt.Println(cc1.name) //ambiguous selector cc1.name
	fmt.Println(cc1.A.name) //ok

	cc2 := C2{} //ok
	fmt.Println(cc2)
	fmt.Println(cc2.name) //ok

}