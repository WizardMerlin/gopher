package main

import "fmt"

type person struct {
	name string
	age int
	//匿名定义嵌套
	contact struct { 
		city, dic string
	}
}

func main() {
	//way1: 定义之后再指定初始化
	pp1 := & person{name:"joe", age:19}
	fmt.Println(*pp1) //{joe 19 { }}
	pp1.contact.city = "上海"
	pp1.contact.dic = "黄浦区"
	fmt.Println(*pp1)

	fmt.Println("--------------")

	//way2: 定义的时候就指定好---失败了
/* 	pp2 := &person {
		name: "mike",
		age: 20,
		contact : {
			city: "上海", dic: "黄浦区",
		},
	} */

	//way2: 定义的时候就制定好，全部直接字面量，不用字段指定---失败了
/* 	pp2 := &person {"mike", 20, {"上海", "黄浦区"}} // missing type in composite literal
	fmt.Println(*pp2) */
}