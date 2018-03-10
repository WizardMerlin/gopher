package main

import "fmt"

func main()  {
	//用标准写法
	s1 := make([]int, 3, 4)
	fmt.Printf("%v, %p\n", s1, s1) //查看地址

	//开始追加
	s1 = append(s1, 1, 2) //现在长度为5已经超过了cap4,所以翻倍为8
	fmt.Println(cap(s1), len(s1))
	fmt.Printf("%v, %p\n", s1, s1) //查看地址(发现已经改变)
	
}