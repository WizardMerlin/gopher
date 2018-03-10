package main

import "fmt"

type empty interface {}

type noempty interface {
    hello()
}

func main() {
    var em empty
    fmt.Println(em == nil) //true

    var no noempty
    fmt.Println(no == nil) //true

    var p *int = nil
    em = p //只要接口收纳了对象, 它就不为空
    fmt.Println(em == nil) //false
}