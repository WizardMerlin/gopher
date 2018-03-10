package main

import "fmt"

func main() {

	i := 5
	switch {
	case 0<=i && i<=5:
		fmt.Println("0-5")
	case 6<=i && i<=10:
		fmt.Println("6-10")
	}

	switch i{
	case 0,1,2,3,4,5:
		fmt.Println("0-5")
	case 6,7,8,9,10:
		fmt.Println("6-10")
	}

	var j int

	switch j, i=1, 1; {
	case i==j :
		fmt.Println("j==b");
		fallthrough
	default:
		fmt.Println("fallthrough 继续执行");
	}
}