package main

import (
	"io"
	"net"
	"fmt"
	"os"
)

func main() {
	//获取命令行参数 (包含要发送的文件名)
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage : ./client filePath")
		return
	}

	filePath := list[1] //发送给 server

	info, err := os.Stat(filePath) //获取文件属性
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//info.Name(), info.Size()

	//连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("dial error : ", err1)
		return
	}
	defer conn.Close()

	//先发送一次文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn Write server err : ", err)
		return
	}

	//坐等接收对方的回应(最好是 "ok" )
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}
	if "ok" == string(buf[:n]) {
		//发送文件
		sendFile(filePath, conn)
	}
}

func sendFile(filePath string, conn net.Conn) {
	//读多少写多少
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF && n ==0 {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err = ", err)
			}
			return 
		}
		conn.Write(buf[:n])
	}

}