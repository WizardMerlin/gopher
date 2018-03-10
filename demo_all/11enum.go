package main

import "fmt"

const (
	a = "A"
	b
	c = iota //c值为2,即取得计数器
	d 		 //d值为3
)

func main()  {
	
	fmt.Println(a) //A
	fmt.Println(b) //A
	fmt.Println(c) //2
	fmt.Println(d) //3

	num := iota //编译报错: undefined: iota
}