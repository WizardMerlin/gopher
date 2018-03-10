package main

import "fmt"

func main() {
	/* 指针数组：数组的元素都是地址 */

	x, y := 1, 2
	//先写好数组，然后把数组的元素类型改为 *int
	array := [2]*int {&x, &y}
	fmt.Println(array); //这样遍历只能拿到地址

	for _,value := range array {
		fmt.Println(*value)
	}
}