package main

import (
	"strings"
	"fmt"
	"bytes"
	"unicode"
)

/*
 重点内容大致上有几个内容:
 1. Querying
 2. Comparing
 3. Trimming
 4. Substituting
 5. Joining
 6. Readers

 并且 strings包 和 bytes包 类似。
*/


func main() {

	//查询: Contains, ContainsAny -- 前者严格匹配，后者存在即可 (包括 Index() 函数)
	fmt.Printf("%v\n", strings.Contains("hello world, gophers", "gophers")) //true
	fmt.Printf("%v\n", bytes.Contains([]byte("abc"), []byte("b"))) //true

	fmt.Printf("%v\n", strings.Contains("hello world, gophers", "gophers ")) //false
	fmt.Printf("%v\n", bytes.Contains([]byte("abc"), []byte("b "))) //false

	fmt.Printf("%v\n", strings.ContainsAny("hello world, gophers", "gophers ")) //true
	fmt.Printf("%v\n", bytes.ContainsAny([]byte("abc"), "b ")) //true
	
	/*-------------------------------------------------------------*/
	fmt.Println()

	//比较: Compare, Equal (string 没有 equal 函数, 比较的话应该使用内建 == > <)
	fmt.Printf("%v\n", "a"=="b")                    //false
	fmt.Printf("%v\n", strings.ToUpper("a") == "A") //true
	fmt.Printf("%v\n", strings.ToUpper("c") == "C") //true

	//bytes 有 Equal() 函数
	fmt.Printf("%v\n", bytes.Equal([]byte{'a', 'b'}, []byte("ab"))) //true

	//bytes 比较使用 Compare
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b'}, []byte("ab"))) // 0
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b'}, []byte("abc"))) // -1
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b', 'c'}, []byte("ab"))) //1

	/*-------------------------------------------------------------*/
	fmt.Println()

	//trim: Trim/TrimFunc/TrimLeft/TrimRight/TrimPrefix/TrimSuffix
	// 最常用的 TrimSpace, 默认 cutset 就是 space characters
	//func Trim(s string, cutset string) string

	fmt.Printf("%q\n",strings.TrimSpace(" space ")) //按照字符串打印
	fmt.Printf("%q\n", bytes.TrimSpace([]byte(" space ")))
	//fmt.Printf("%v\n", bytes.TrimSpace([]byte(" space ")))

	//普通的Trim函数 (只能 trim prefix/postfix)
	fmt.Print(strings.Trim("¡¡¡Hello, !¡Gophers!!!\n", "!¡")) //Hello, Gophers

	//指定自定义的 trim 内容 (只能 trim prefix/postfix)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) //不是 '字符'的全部被 trim
	}
	fmt.Printf("%q\n", strings.TrimFunc("  123hello 234 !word\n", f)) //hello 234 !word

	//替换: replace (中间的也内容如果要被替换，应该使用 replace)
	//func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // <0 表示全部替换

	//使用 Replacer 指定 old, new 串 (定制 replace)
	replacer := strings.NewReplacer("alpha", "a", "beta", "b")
	fmt.Printf("%q\n", replacer.Replace("alpha beta, beta alpha, over"))

	/*-------------------------------------------------------------*/
	fmt.Println()


	//子串: substituting (Fields / Split) + Join

	// Count
	fmt.Println(strings.Count("cheese", "e")) //3
	fmt.Println(strings.Count("five", "")) // 子串是空串，返回: 字符个数 + 1

	//Fields 按照空白字符进行分隔
	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz  "))

	//join
	str := " foo bar baz "
	fields := strings.Fields(str)
	fmt.Printf("%q\n", strings.Join(fields, "|"))

	//split 分隔 (可以指定具体的分隔符) 还有 SplitN 指定分隔为几部分, SplitAfterN, SplitAfter
	fmt.Printf("%q\n", strings.Split(strings.Join(fields, "|"), "|"))
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	/*-------------------------------------------------------------*/
	fmt.Println()


	//Reader: reading from a string (从字符串里读)
	//(这个类型实现了多个读接口; 凡是需要这些接口的内容也可以使用这个 string 的包装类型)
	var a, b, c string
	s := "a b c"
	reader := strings.NewReader(s) // string 包装类
	if _, err := fmt.Fscan(reader, &a, &b, &c); err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("a:%q, b:%q, c:%q\n", a, b, c)

	/*-------------------------------------------------------------*/
	fmt.Println()

	//Builder : 有 reader 当然也有 wrieter (往字符串里写, 高效的写)
	//builder 其实是对 writer 的一个封装(因为 string 不可变, 只能是拷贝副本写入；这里的Builder则不是用副本，所以高效)
	//go 1.10 版本， go 1.9版本不支持
	var writer strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&writer, "%d...", i)
	}
	writer.WriteString("ignition")
	fmt.Println(writer.String())

}