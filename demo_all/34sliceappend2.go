package main

import "fmt"

func main() {
	arr := []int{1,2,3}
	s1 := arr[0:2]
	s2 := arr[1:3]

	fmt.Println(s1) //1, 2
	fmt.Println(s2) //2, 3

	//s1和s2共享一个 2 元素，修改试试看(同时改变)
	s2[0] = 4
	fmt.Println(s1) //1, 4
	fmt.Println(s2) //4, 5

	//看看s2的cap
	fmt.Println(cap(s2))

	//但是在 s2 背后添加呢?
	s2 = append(s2, 2, 3, 4) //重新分配然后拷贝过去, 底层数组已经不一样了
	fmt.Println(s2)

	//在修改s2已经影响不了s1了
	s2[0] = 9
	fmt.Println(s1) //1, 4
	fmt.Println(s2) //9, 3, 2, 3,4
}