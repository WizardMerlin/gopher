package main

import "fmt"

func main()  {
	a := 65
	str := string(a)

	fmt.Println(str)//打印字符A, 而不是65
}