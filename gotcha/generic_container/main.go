package main

import (
	"fmt"
)

//创建一个 generic container, 底层用 slice
type Container []interface{}

func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem) //append 方法可能比较危险(重新分配底层数组问题)，所以重新赋值一下
}
func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem //返回顶元素
}
func (c *Container) IsEmpty() bool {
	return len(*c) == 0
}

type IntContainer struct {
	Container
}

func (ic *IntContainer) Put(elem int) { //入参 具体化
	(*ic).Container.Put(elem)
}

func main() {
/* 	intContainer := &Container{}
	intContainer.Put(1)
	//intContainer.Put("5") //Get() did not return an integer.
	intContainer.Put(3) */

	intContainer := &IntContainer{Container{}}
	intContainer.Put(1)
	//intContainer.Put("5") //只要你这么写，编译时就报错了，不用等到运行时了
	intContainer.Put(3)



	for !intContainer.IsEmpty() {
		elem, ok := intContainer.Get().(int)
		if !ok {
			fmt.Println("Get() did not return an integer.")
			continue
		}
		fmt.Printf("Elem: %d (%[1]T)\n", elem)
	}

}