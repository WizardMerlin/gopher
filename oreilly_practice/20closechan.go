package main

import ( 
	"fmt"
	"time"
)

func printer(msg string, goChan chan bool) {
	<- goChan //阻塞等待读

	fmt.Printf("%s\n", msg)
}

func main() {
	goChan := make(chan bool)

	for i:= 0; i< 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), goChan)
	}
	
	time.Sleep(5 * time.Second)
	close(goChan) //5秒过，那些阻塞等待的goroutines可以继续执行
	time.Sleep(5 * time.Second)
}