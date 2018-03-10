package main

import "fmt"

func main() {
	//操作副本不会影响原来的对象
	sli := make([]map[int]string, 3, 4) //长度为3, cap为4的slice
	//遍历slice进行初始化
	for _,value := range sli {
		value = make(map[int]string, 1)
		value[1]="OK"
		fmt.Println(value)
	}
	fmt.Println(sli) //这里是空的，表明是对副本操作了。

	//直接操作原始 sli
	//遍历slice进行初始化
	for i := range sli {
		sli[i] = make(map[int]string, 1)
		sli[i][1]="OK"
		fmt.Println(sli[i])
	}
	fmt.Println(sli)

}