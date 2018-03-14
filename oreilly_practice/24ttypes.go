package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

type webPage struct {
	url string
	body []byte
	err error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		fmt.Println("get Page err")
		w.err = err
		return
	}
	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func main() {
	w := &webPage{url: "http://www.baidu.com/"}
	w.get()
	fmt.Printf("url: %s, %d length; error? %v\n", w.url, len(w.body), w.err)
}