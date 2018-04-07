package main

import (
	"io"
	"strings"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("err = ", err)
	}
	defer listener.Close()

	//并发接收多个请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("err = ", err)
		}
		//为每一个连接开一个协程
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取对端地址
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + "addr connect succeed")

	//读取用户请求
	buf := make([]byte, 1024*4)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("err = ", err)
			}
			return
		}
		fmt.Println("read buf = ", string(buf[:n])) //查看读取到的内容
	
		conn.Write([]byte(strings.ToUpper(string(buf[:n])))) //转成大写返回给客户端
	}

}