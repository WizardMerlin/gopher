package main

import "fmt"
import "sort"

func main() {
	m1 := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"} 
	//初始化的时候貌似有序，但是遍历时发现是无序的
	for k,v := range m1 {
		fmt.Println(k,v) //多运行几次，发现是无序的
	}

	//此时可以把 key 放入容器中，对 key 排序，从而拿到对 value 排序的目的
	//意思是，如果slice有序，那么map一定有序
	i := 0
	sli := make([]int, 5, 5)
	for k := range m1 {
		sli[i] = k
		i++
	}
	fmt.Println(sli)
	sort.Ints(sli) //排序 int 类型的容器
	fmt.Println(sli)
}