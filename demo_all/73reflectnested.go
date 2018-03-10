package main

import (
	"fmt"
	"reflect"
)


type User struct {
	Id int
	Name string
	Age int
}

type Manager struct {
	User //只写了类型, 相当于匿名字段
	title string
}

func main() {
	m := Manager{User{1, "bob", 30}, "manager"}
	
	//反射代码 - 对类型操作
	typeManger := reflect.TypeOf(m)

	//Field 拿到完整的匿名结构
	fmt.Printf("%v\n", typeManger.Field(0)) 
	fmt.Printf("%#v\n", typeManger.Field(0)) //输出完整结构信息

	//FieldByIndex 拿到匿名结构内的所有字段, 以 slice 索引的形式
	//直接拿到 manager, 相当于 Field(0)
	fmt.Printf("%v\n", typeManger.FieldByIndex([]int{0}))
	//取得 user.id 
	fmt.Printf("%v\n", typeManger.FieldByIndex([]int{0, 0})) //直接传递的 slice 要指定两级索引值
	//相对于 Manager 而言 User 是第0个，相对于 User 而言 Id 是第0个。
	//取得 user.name
	fmt.Printf("%v\n", typeManger.FieldByIndex([]int{0, 1}))

	//取得 manager.title
	fmt.Printf("%v\n", typeManger.FieldByIndex([]int{1}))

	//FieldByName, 同样用于取得 manager.title
	titleName,_ := typeManger.FieldByName("title")
	fmt.Printf("%v\n", titleName)
}