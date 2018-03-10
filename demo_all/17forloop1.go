package main

import "fmt"

func main() {
	
	b := 2

	for {			
		if b == 2 {
			fmt.Println("b==2, 第一次循环");
			b--
		} else if b ==1 {
			fmt.Println("b==1, 直接 continue 吧");
			b--
			continue
		} else {
			fmt.Println("b==0, 顺序执行下去算了")
		}

		if b==0 {
			fmt.Println("已经执行到最后流程了，出去吧")
			break
		}

	}
}