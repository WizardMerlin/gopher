package main

import "fmt"

func main()  {
	a := 1
	//b := a++ //编译报错: syntax error: unexpected ++ at end of statement
	//b := ++a //编译报错: syntax error:unexpected ++, expecting expression

	//++a //报错: unexpected ++, expecting }, 不能放在变量的左边
	a++ //只能放在变量的右边
	b := a

	fmt.Println(b)
}