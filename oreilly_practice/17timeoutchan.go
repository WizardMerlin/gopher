package main

import (
	"time"
	"fmt"
)

func emit(wordChan chan string, done chan bool) {
	words := []string{"the","quick","brown","fox"}

	i := 0
	t := time.NewTimer(3 * time.Second) //三秒后执行 timeout 分支

	defer close(wordChan)

	for {
		select {
			case wordChan <- words[i]: //写出去
			i++
			if i == len(words) {
				i = 0 //再重头开始
			}
		case <-done: //收到停止的信息，停止接收
			fmt.Println("Got done")
			//close(done)
			done <- true
			return
		case <-t.C:
			fmt.Println("tiemout")
			return
		}
	}
}

func main() {
	wordChan := make(chan string)
	doneChan := make(chan bool)

	go emit(wordChan, doneChan)

	for word := range wordChan {
		fmt.Println(word)
	}
}