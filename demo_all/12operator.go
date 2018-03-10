package main

import "fmt"

func main()  {
	//取反
	fmt.Println(!true) //false
	//fmt.Println(!1) //编译报错: ! untyped number

	//^当一元运算符
	fmt.Print("^一元运算: ")
	fmt.Println(^1) //-2 结果比较奇怪, 用的比较少



	/* 位运算符 */
	//与运算符
	fmt.Println(1&0) //0
	//或运算
	fmt.Println(1|0) //1

	//亦或
	fmt.Println(1^0) //1
	fmt.Println(0^1) //1
	fmt.Println(1^1) //0
	fmt.Println(0^0) //0
	//fmt.Println(false^false); //编译报错:^ not defined on bool

	fmt.Println("-----special-----")
	//&^ 
	//主要看第二个操作数:
	//如果第二个操作数是0，那么结果和第一个操作数的位相同；
	//第二个操作数的位是1，那么结果一定是0
	fmt.Println(1&^0) //1
	fmt.Println(0&^1) //0
	fmt.Println(0&^0) //0
	fmt.Println(0&^1) //0
	fmt.Println(1&^1) //0

	fmt.Println("-----special-----")

	//乘除，取余-略

	// << >> 
	fmt.Println(1<<10) // 1024
	fmt.Println(8>>3) //1


	//逻辑运算符 (带短路)
	fmt.Println(false && true)  //false
	fmt.Println(true && false)  //false
	fmt.Println(false || true)  //true

}