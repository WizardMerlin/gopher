package main

import "fmt"
import "reflect"

type User struct {
	Id int
	Name string
	Age int
}

type empty interface{}

func (usr User) Print() {
	fmt.Printf(usr.Name, usr.Id, usr.Age)
}

func main() {
	uu := User {1,"bob",20}
	//uu.Info(&uu) //当你传递了指针，则报错
	uu.Info(uu)
}

//info 方法内部用到了反射
func (usr User)Info(obj empty) {
	//拿到类型
	t := reflect.TypeOf(obj) // t 是 Type接口的实例副本
	fmt.Println("Type:", t.Name()) //结构类型是有方法的，所以可以调用 Name()

	if kind := t.Kind(); kind != reflect.Struct {
		fmt.Println("传入反射函数的参数不是实例类型，您是不是传入了指针？")
		return
	}

	//拿到值
	v := reflect.ValueOf(obj) //这里拿到的结构 (而不是字段)
	//字段循环取 (字段的类型、名字可以通过 t 获取；但是 value 只能通过 v 获取)
	for i := 0; i < t.NumField(); i++ {
		//拿到字段的名称，类型 (通过t)
		field := t.Field(i) //拿到字段对象 (可以查看字段的类型，名称；但是拿不到值)
		fmt.Print(field.Name," ", field.Type," ") //字段没有方法，直接用属性

		//拿到字段的值
		val := v.Field(i).Interface() //拿到 v当前值
		fmt.Println(val)
	}

	//调用 obj 的方法 (可能有多个方法，所以也是循环)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type) //Type表示方法的签名 //Print: func(main.User)
	}
}

/*
Type: User
Id int 1
Name string bob
Age int 20
*/