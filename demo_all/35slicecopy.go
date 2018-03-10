package main

import "fmt"

func main() {
	s1 := []int{1,2,3,4}
	s2 := []int{5,6,7}

	//copy(s1,s2) //s2拷贝到s1
	//fmt.Println(s1)  //s1 变成 5,6,7,4

	//copy(s2, s1) //s1拷贝到s2 (原来的 slice 长度保持不变, 所以只拷贝3个元素)
	//fmt.Println(s2) //1,2,3

	//选择性拷贝
	copy(s2, s1[1:]) //2,3,4 拷贝到 s2 覆盖
	fmt.Println(s2)

	//完全拷贝？请直接赋值
	s3 := s1
	fmt.Println(s3)
}