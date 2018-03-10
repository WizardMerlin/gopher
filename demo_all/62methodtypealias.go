package main

import "fmt"

//类型别名，然后给别名指定方法
type XX int

func main() {
	var x XX
	x.printInt()
	x.setInt(3)
}

//给 XX 指定方法

func (p *XX)printInt() {
	*p = *p + 1 //因为 X 就是 int
	fmt.Println(*p)
}

func (p *XX)setInt(num int) {
	//*p += num //编译报错 // mismatched types XX and int
/* 	tmp := XX(num)
	*p += tmp  */
	*p += XX(num)
	fmt.Println(*p) 
}