package main

import "fmt"

//常量
const defalt_salary int = 10000

//默认常量
const default_age = 20

//常量组
const (
	name = "bob"
	age = 22
	salary = defalt_salary
	married  //如果不指定，则是上一行的表达式
)

//多赋值
const a, b = 1, 1.1


//套用上一行 + 多赋值
const (
	c, d = 2, "2,2"
	e, f
)

func main()  {
	fmt.Println(married) //输出 10000
	fmt.Println(a) // 1
	fmt.Println(b) // 1.1

	fmt.Println(e) //2
	fmt.Println(f) //2.2
}