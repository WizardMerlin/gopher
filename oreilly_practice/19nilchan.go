package main

import (
	"time"
	"fmt"
	"math/rand"
)

func reader(ch chan int) {
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-t.C: //3秒钟到了，暂停接收(ignore, 忽略)--chan接收一直走该分支
			ch = nil 
		}
	}
}

func writer(ch chan int) {
	t := time.NewTimer(2 * time.Second)
	for {
		select {
		case ch<- rand.Intn(10):
		case <- t.C:
			ch = nil
		}
	}
}

func main() {

	ch := make(chan int)
	go reader(ch)
	go writer(ch)

	time.Sleep(10 * time.Second)
}
