package main

import "fmt"

func main()  {
	//var a int = 1
	a := 1
	//var b *int = &a
	b := &a

	fmt.Println(b); //打印地址
	fmt.Println(*b); //打印值

	var p *int
	fmt.Println(p) //<nil>
}