package main

import "fmt"

func main() {
	arr := [5]int{2,8,4,9,5}
	fmt.Println(arr)  //排序之前

	//降序排序吧
	leng := len(arr)
	for i := 0; i < leng; i++ {
		for j := i+1; j < leng; j++ {
			if arr[i] < arr[j] {
				//交换 保证 当前的 arr[j]是最小
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		} //inner for
	} //outer for
	fmt.Println(arr)
}