package main

import "fmt"

func emit(c chan string) {
	words := []string {"the", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word //循环写入 chan
	}
	close(c)
}

func main() {
	wordChan := make(chan string) //不指定队列大小，默认不使用队列

	go emit(wordChan)

	for word := range wordChan { //事先并不知知道要对方要写几个
		fmt.Printf("%s ", word)
	}

	//word := <- wordChan  //这样只能拿到一个, 然后程序结束

	fmt.Print("\n")
}