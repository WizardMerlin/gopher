package main

import (
	"strconv"
	"fmt"
)

func main()  {
	var a = 65
	str := strconv.Itoa(a)
	//a = strconv.Atoi(b)
	fmt.Println(str)//65
}