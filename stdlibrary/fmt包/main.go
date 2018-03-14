package main

import (
	"log"
	"fmt"
)

type point struct {
	x,y int64
}

func main() {
	fmt.Println("主要练习 fmt 包下面的 print相关，scan相关；自定义的 formatter, scanner")

	//scanf
	p1 := point{}
	fmt.Println(p1.x, p1.y)
	if _, err := fmt.Sscan("5 6", &p1.x, &p1.y); err != nil {
		fmt.Println(err)
		log.Fatal("Sscan error")
	}
	fmt.Println(p1.x, p1.y)

	p2 := point{}
	fmt.Println(p2.x, p2.y)
	if _, err := fmt.Scan(&p2.x, &p2.y); err != nil {
		fmt.Println(err)
		log.Fatal("Scan error") //输入字符也会报错
	}
	fmt.Println(p2.x, p2.y)

	p3 := point{}
	fmt.Println(p3.x, p3.y)
	if _, err:= fmt.Scanf("%d,%d", &p3.x, &p3.y); err != nil { //对输入格式有严格要求
		fmt.Println(err)
		log.Fatal("Scanf error")
	}
	fmt.Println(p3.x, p3.y)


	//String(), GoString()
	//fmt.Printf("%v\n", p1)
	//fmt.Printf("%#v\n", p2)

	//Scanner 自定义
	p4 := point{}
	if _, err := fmt.Scanf("%$", &p4); err != nil { //自动调用 Sanner 接口指定的输入方式
		fmt.Println(err)
		log.Fatal("Scanner error")
	}
	fmt.Printf("%v\n", p4)
	fmt.Printf("%$", p4)


}

func (p point) String() string {
	fmt.Println("calling String()")
	return fmt.Sprintf("Pointer(%d, %d)", p.x, p.y)
}

func (p point) GoString() string {
	fmt.Println("calling GoString()")
	return fmt.Sprintf("Point(%d, %d)", p.x, p.y)
}

// "%$" input
func (p *point) Scan(f fmt.ScanState, v rune) error {
	fmt.Println("calling Scan() to input")
	switch v {
	case '$':
		n, err := fmt.Fscanf(f, "%v,%v", &p.x, &p.y)
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("2 values expected, got %d", n)
		}
	}
	return nil
}

// "%$" output --- formatter
func (p point) Format(f fmt.State, v rune) {
	switch v {
	case '$':
		fmt.Fprintf(f, "%v,%v\n", p.x, p.y)
	case 'v':
		//fmt.Printf("%v\n", p) //会调用 Format 造成死循环。
		fmt.Print("calling Format; 之后不再调用 String()")
	default:
		fmt.Fprintf(f, "point(%v,%v)\n", p.x, p.y)
	}
}