package main

import "fmt"
import "reflect"

type User struct {
	id int
	name string
	age int
}

type empty interface{}

func (usr User) print() {
	fmt.Printf(usr.name, usr.id, usr.age)
}

func main() {
	uu := User {1,"bob",20}
	uu.info(uu)
}

//info 方法内部用到了反射
func (usr User)info(obj empty) {
	//拿到类型
	t := reflect.TypeOf(obj) // t 是 Type接口的实例副本
	fmt.Println("Type:", t.Name()) //结构类型是有方法的，所以可以调用 Name()

	//拿到值
	v := reflect.ValueOf(obj) //这里拿到的结构 (而不是字段)
	//字段循环取 (字段的类型、名字可以通过 t 获取；但是 value 只能通过 v 获取)
	for i := 0; i < t.NumField(); i++ {
		//拿到字段的名称，类型 (通过t)
		field := t.Field(i) //拿到字段对象 (可以查看字段的类型，名称；但是拿不到值)
		fmt.Print(field.Name, field.Type) //字段没有方法，直接用属性

		//拿到字段的值
		val := v.Field(i).Interface() //拿到 v当前值
		fmt.Println(val)
	}
}
//编译报错了
// cannot return value obtained from unexported field or method
