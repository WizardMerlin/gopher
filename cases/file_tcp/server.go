package main

import (
	"io"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	defer listener.Close()

	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err = ", err1)
		return
	}
	defer conn.Close()

	//读取客户端请求
	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn Read err = ", err1)
		return
	}

	fileName := string(buf[:n])
	conn.Write([]byte("ok"))

	//介绍文件内容
	recvFile(fileName, conn)
}

func recvFile(fileName string, conn net.Conn) {
	//新建一个文件
	f ,err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file err = ", err)
		return
	}
	defer f.Close()

	//从 conn 读取然后写入本地文件
	buf := make([]byte, 1024*4)

	//接收多少写多少
	for{
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("conn.Read err = ", err)
		}
		if n == 0 {
			fmt.Println("文件接收完毕")
			return
		}
		f.Write(buf[:n])
	}
}