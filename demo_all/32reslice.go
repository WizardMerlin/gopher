package main

import "fmt"

func main() {
	oldSlice := []byte{'a','b','c','d','e'} //5个元素
	sliceA := oldSlice[0:3] //取a, b ,c
	fmt.Println(string(sliceA))
	sliceB := oldSlice[3:5] //取 d,e
	fmt.Println(string(sliceB))

	//sliceC 是在 sliceB 的基础上取
	sliceC := sliceB[1:] //取e
	fmt.Println(string(sliceC))

	//看看 len 和 cap
	fmt.Println(len(oldSlice), cap(oldSlice)) //5 5
	fmt.Println(len(sliceA), cap(sliceA)) //3, 5
	fmt.Println(len(sliceB), cap(sliceB)) //2, 2
	fmt.Println(len(sliceC), cap(sliceC)) //1, 1
}