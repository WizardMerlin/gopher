package main

import "fmt"

func main() {
	b := 1
	b = 2
	defer subFunc(b) //0xc4200120b0
	b = 3
	fmt.Printf("outer b: %v, %p\n", b, &b) //0xc420012090
}


func subFunc(b int) {
	fmt.Printf("defer b: %v, %p\n", b, &b)
}

/*运行结果: (通过参数传递的，所以地址不同；defer时拿到的是一份拷贝)
outer b: 3, 0xc420012090
defer b: 2, 0xc4200120b0
 */