package main

import "fmt"

type human struct {
	name string
	age int
}

type student struct {
	human //非匿名嵌套
	stuNo int
}

type teacher struct {
	human
	teaNo int
}

func main() {

	//定义就初始化--直接字面量(这也解释了为啥内部匿名嵌套定义不能这么初始化)
	ss := student{human{"bob",10}, 1} //ok
	fmt.Println(ss)

	//定制就初始化--按照字段指定
	ss2 := student {
		human:  human{
					name: "bob",
					age: 10,
				}, 
		stuNo: 1,
	}
	fmt.Println(ss2)

	//操作嵌入式结构的字段
	ss2.human.age = 11
	fmt.Println(ss2)

	//相当于把嵌入结构字段----不推荐这样，容易引起歧义
	ss2.age = 12
	fmt.Println(ss2)
}