package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 主要在基本类型(int, float, bool等)，string类型，quote方法
	// 常用的 Atoi, Itoa; 以及 Parse 方法, Format 方法
	
	s := "1604"
	i, _ := strconv.ParseInt(s, 8, 0) //8进制，bitSize 0 则表示默认是机器字长
	fmt.Printf("%v - %T\n", i, i)

	s1 := strconv.Itoa(35) // 包装的 Format 方法
	fmt.Printf("%v - %T\n", s1, s1)

	b := int64(-1024)
	t := strconv.FormatInt(b, 2)
	fmt.Printf("%T, %v\n", t, t)

	//append方法
	bb := []byte("what are we appending ?:")
	bb = strconv.AppendInt(bb, -345, 10) // golang 1.10版本
	fmt.Printf("%v\n", string(bb))

	//float 转换
	π := "3.1415926535"
	f, err := strconv.ParseFloat(π, 64)
	fmt.Printf("%v - %T -%v\n", f, f, err)

	//指定底和精度
	π1 := 3.1415926535
	f1 := strconv.FormatFloat(π1, 'e', 2, 64) //float64, 精度为2
	fmt.Printf("%v - %T\n", f1, f1)

	//bool 类型
	list := []string{"1", "t", "T", "TRUE", "true", "Yes",
		 "Yay", "Yippie", "0", "f", "false", "False", "No"}
	for _, bs := range list {
		bs, err := strconv.ParseBool(bs)
		fmt.Printf("%v - %T - %v\n", bs, bs, err)
	}

	//Quote 转换 (不能转换成 ascii码的，会以 /u开头)
	fmt.Println(strconv.Quote("There is  a π symbol here as well as a tab	."))
	fmt.Println(strconv.QuoteToASCII("There is  a π symbol here as well as a tab	."))


	fmt.Println()

	// 还有一个 strconv Error 处理 (parse 标准写法)
	str := "Not a number"
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		e := err.(*strconv.NumError)
		fmt.Println("Func:", e.Func)
		fmt.Println("Num:", e.Num) // shows the input
		fmt.Println("Err:", e.Err) //具体的 error
		fmt.Println(err)
	}


}