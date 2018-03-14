package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("te.txt")
	if err != nil {
		fmt.Println("Open File Error")
		os.Exit(1)
	}
	defer f.Close()

	bytes := make([]byte, 128)
	n, err := f.Read(bytes)
	str := string(bytes) //编译时检查
	fmt.Println("read ", str, " : ", n, " bytes")
}