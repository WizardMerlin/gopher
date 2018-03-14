package main

import (
	"time"
	"fmt"
	"math/rand"
	"sync/atomic"
)

var (
	runningCount int64 = 0
)

func worker(sema chan bool) {
	//<- sema //队列为空则等待
	atomic.AddInt64(&runningCount, 1) //消耗一个资源

	//假装在做任务
	fmt.Printf("[ %d", runningCount)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Print("]")

	//完成任务了
	//sema <-true
	atomic.AddInt64(&runningCount, -1) //把资源还回来
}

func main() {

	sema := make(chan bool, 10)

	//开启 1000 个协程
	for i:= 0; i < 1000; i++ {
		go worker(sema)
	}

	for i:= 0; i< 10; i++ {
		sema <-true //虽然开启了 1000 个协程，但实际工作的只有 10个
	}

	time.Sleep(10 * time.Second) //让子 routines 有足够的时间执行
}