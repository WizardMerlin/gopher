package main

import "fmt"

func main() {
	var sla []int //不指定长度即是 slice
	fmt.Println(sla)

	//已经存在一个数组了, 用slice截取后面的元素
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr)
	//sli := arr[1] //取一个元素
	
	//sli := arr[5:10] //含头不含尾 index 5-9, index 10 不包含
	//等价于//sli := arr[5:len(arr)]
	sli := arr[5:] //也是取 index 5-9, 直接取到末尾
	fmt.Println(sli)
}