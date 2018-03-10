package main

import "fmt"

func main() {
	//简单演示一下 range for

	arr := [5]int{1,2,3,4,5}

	//使用 range 关键字，后面跟着容器的名字
	//可以把 range 当做一个多返回值的函数
	/* 
	for key, value := range oldMap {
    	newMap[key] = value
	} */
	for index,value := range arr {
		fmt.Printf("index=%d, value=%d\n", index, value);
	} 
}