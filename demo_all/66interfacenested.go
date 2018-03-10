package main

import "fmt"

type A interface {
	getName() string
	xxx() 
	yyy()
}

type B interface {
	getName() string //返回值不能做作为区分
	zzz()
	ssss()
}

/*duplicate method getName*/
type C1 interface {
	A
	B
}

type CC1 struct {
	name string
}

func main() {
	//var c C1
	c := CC1{"ccc"}
	//初步结论，go的设计团队没有在语言层面上解决这个问题，只能做一些检查。
	fmt.Println((*CC1).getName(&c)) //编译报错:ambiguous selector c.getName
}

func (c CC1) getName() string{
	return c.name 
}

/* type C2 interface {
	A
	B
	name() string
}
 */