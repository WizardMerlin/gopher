package main

import "fmt"

func main() {
LABEL1:
	for { //外层是个死循环
		for i:=0; i < 10; i++ {
			if i==1 {
				break LABEL1;
			}
		}
	}
	fmt.Println("跳出了与 LABEL 同级的循环")
}