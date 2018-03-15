package main

import (
	"fmt"
	"regexp"
)

func main() {
	//主要用途还是编译过后，拿到编译对象，然后使用其成员方法， match, find, index, replace, split.

	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`) //编译正则表达式

	fmt.Println(validID.MatchString("adam[23]")) //true
	fmt.Println(validID.MatchString("eve[7]"))   //true
	fmt.Println(validID.MatchString("Job[48]"))  //false
	fmt.Println(validID.MatchString("snakey"))   //false
}