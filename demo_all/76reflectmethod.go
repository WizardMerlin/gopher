package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id int
	name string
	age int
}

//给它添加一个方法

func (u *User) SetUserAge(age int) {
	//fmt.Println("param age = ", age)
	u.age = age
}

func (u User) Print() {
	fmt.Println(u.name, " : ", u.age)
}

func main() {
	u := &User{1, "bob", 20}
	fmt.Println(u.age) //20
	u.SetUserAge(21)
	fmt.Println(u.age) //21 正常调用 ok

	//反射调用 (先拿到反射 Value 对象)
	valObj := reflect.ValueOf(u) //传入的是指针
	
	/*应该做一些检查*/

	//拿到方法
	m := valObj.MethodByName("SetUserAge")
	//设置参数
	args := []reflect.Value{reflect.ValueOf(22)} //int类型的 Value

	//传入参数，反射调用
	m.Call(args)
	fmt.Println(u.age) //22

	m1 := valObj.MethodByName("Print")
	m1.Call([]reflect.Value{}) //无参数调用
}