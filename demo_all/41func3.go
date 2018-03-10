package main

import "fmt"

func main() {

	a := 1
	b := &a
	c := []int{1,1}
	

	f1(a)
	fmt.Println("外部a:", a)


	f2(b)
	fmt.Println("外部b:", *b) //被内部修改

	f3(c)
	fmt.Println("外部c:", c) //被内部修改
}

func f1(a int) {
	a = 0
	fmt.Println("内部a", a)
}

func f2(b *int) {
	*b = 0
	fmt.Println("内部b", *b)
}

func f3(c []int) { //传递slice
	c[0] = 0
	c[1] = 0
	fmt.Println("内部c", c)
}