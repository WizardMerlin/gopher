package main

import "fmt"

func main() {
	a := [2]int{1,2}
	b := [2]int{1,2}
	c := [2]int{2,3}
	//d := [1]int{}

	fmt.Println(a==b) //true
	fmt.Println(a==c) //false
	//fmt.Println(a==d) //不同类型的比较，编译报错
	//invalid operation: a == d (mismatched types [2]int and [1]int)

	aa := [...]string{"hello"}
	bb := [...]string{"hello"}
	//cc := [...]string{"hello", "world"}

	fmt.Println(aa==bb) //true
	//fmt.Println(aa==cc) //不同类型的比较，编译报错
	//invalid operation: aa == cc (mismatched types [1]string and [2]string)
}