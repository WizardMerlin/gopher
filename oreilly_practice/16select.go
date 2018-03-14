package main

import "fmt"

func emit(wordChan chan string, done chan bool) {
	words := []string{"the","quick","brown","fox"}

	i := 0

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
		}
	}
}

func main() {
	wordChan := make(chan string)
	doneChan := make(chan bool)

	go emit(wordChan, doneChan)

	for i:= 0; i < 10; i++ {
		fmt.Printf("%s\n", <-wordChan) //收取对方写入的内容，10次
	}
	doneChan <- true
	<- doneChan  //等着接收子 routine 发送最后的确认给我
}