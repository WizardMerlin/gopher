package main

import (
	"bytes"
	"fmt"
)

// Map 函数可以逐个字符替换或者处理 Slice 的元素

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}

		if r == '3' {
			return '4'
		}

		return r
	}
	fmt.Printf("%s\n", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher 32...")))
}