package main

import "fmt"

func main() {
	//多维数组一样可以使用下标运算, 索引指定等
	//只有第一维可以用 ... 指定长度(自动计算)，最好自己指定

	arr := [...][3]int {
		{1,2,3},
		{1}}
	//}//编译报错

	fmt.Println(arr)
}