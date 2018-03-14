package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPage(url string)(int, error) { //返回读取的长度
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		return 0, err
	}
	return len(body),nil
}

func getter(url string, size chan int) {
	pageLen, err := getPage(url)
	if err == nil {
		size <- pageLen
	}
}

func main() {
	urls := []string{ "http://www.baidu.com/", "http://www.merlinblog.site/",
		 "http://www.commoncommonheart.com/"}

	sizeChan := make(chan int)

	for _, url := range urls {
		go getter(url, sizeChan)
	}

	//等待接收
	for i:=0; i<len(urls); i++ {
		fmt.Printf("%s is length %d\n", urls[i], <-sizeChan)
	}


}
