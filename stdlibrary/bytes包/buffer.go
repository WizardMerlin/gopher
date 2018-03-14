package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello ")) //往 Slice 里面写
	fmt.Fprintf(&b, "world!") //往 Slice 里面写
	b.WriteTo(os.Stdout)
}
