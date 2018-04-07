package main

import (
	"os"
	"fmt"
	"net"
)

/*直接发送内容*/
/* func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("big")) //发送数据
} */

//从 os.Stdin读取内容，然后在发送

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("dial err : ", err)
		return
	}
	defer conn.Close()

	go func() {
		//从键盘设备文件读
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)  //读取标准输入
			if err != nil {
				fmt.Println("os.Stdin, err = ", err)
				return
			}

			//把输入的内容给服务器发送
			conn.Write(str[:n]) //读多少写多少
		}

	}()

	//主协程负责死循环在这守着服务器的回写 
	//(好处是服务器关闭，这边儿可以退出; 否则放到子协程中读，主协程进行 stdin输出，
		//则即使读到服务器关闭的异常，程序也无法退出)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务器的写回
		if err != nil {
			fmt.Println("read err = ", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印服务器拿过来的数据
	}
}