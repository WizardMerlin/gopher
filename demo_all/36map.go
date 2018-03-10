package main

import "fmt"

func main() {
	//全写
	var m1 map[int]string //定义
	m1 = map[int]string{} //初始化
	fmt.Println(m1) //map[]

	//或者使用 make
	var m2 map[int]string = make(map[int]string, 3) //cap 可以省略
	fmt.Println(m2)

	//make也可以简写 (系统也建议你简写)
	m3 := make(map[int]string)
	fmt.Println(m3)

	/*添加元素*/
	m1[1] = "Ok"
	fmt.Println(m1[1])

	/*覆盖, 重复添加的话，就直接覆盖了*/ 
	m1[1]="Cancel"
	fmt.Println(m1[1])

	/*删除, by key*/
	delete(m1, 1)
	fmt.Println(m1[1]) //打印为空

}