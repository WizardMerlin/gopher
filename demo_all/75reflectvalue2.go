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

func main() {
	u := &User{0, "bob", 20}
	fmt.Println(*u)

	Modify(u) //传入指针
	fmt.Println(*u)
}


func Modify(o interface{}) { 

	//先拿到反射对象 reflect.Value 类型
	valObj := reflect.ValueOf(o) 
	//var val Value  //没必要这么定义，valObj本身类型就是 Value

	//传入空接口一定要检查入参
	if valObj.Kind() != reflect.Ptr || !valObj.Elem().CanSet() {
		fmt.Println("传入不是指针或者不能改修原结构, 下面操作无法完成，直接退出")
		return
	} else {
		valObj = valObj.Elem() //类型都是 Value 可以直接赋值
	}

/* 	if age := valObj.FieldByName("Age"); age.Kind() ==reflect.Int {
		age.SetInt(21)
	} */

	age := valObj.FieldByName("Age")

	if(!age.IsValid()) { //没有找到返回的是 []Value{}
		fmt.Println("没有找到相关的字段")
		return
	} //实际上没有找到字段去设置也是不成功的
	if age.Kind() == reflect.Int {
		age.SetInt(22)
	}

}