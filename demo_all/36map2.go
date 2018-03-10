package main

import "fmt"

func main() {
	complex_map := make(map[int]map[int]string) //value的类型是 map[int]string
	
	//先取一下值，看看存不存在
	str, ok := complex_map[1][1]
	if !ok {
		//说明第一个key还没有插入，即没有值; 那就创建一下
		complex_map[1] = make(map[int]string)
		//现在可以插入了
		complex_map[1][1] = "Ok"
	}

	//取出来看看
	str, ok = complex_map[1][1]
	fmt.Println(str, ok)
}