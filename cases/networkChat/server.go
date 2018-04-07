package main

import (
	"time"
	"strings"
	"net"
	"log"
)

//用于记录全局在线人员
type Client struct {
	C chan string //用于给该Client转发消息
	Name string  //用户名
	Addr string  //网络地址
}

var onlineMap map[string]Client

var message = make(chan string) //用于操作 onlineMap 同步控制，同时传输消息


//总体流程
//server -(处理连接 或者 转发)-> message --->Client.C ---> 读conn.Write给Client端
func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("net.Listen = ", err)
	}
	defer listener.Close()

	//初始化 map
	onlineMap = make(map[string]Client)


	//协程1，用于转发消息给 Client.C
	go sendMessage2ClientChan()

	//主协程，循环等待连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listen.Accept err = ", err)
			continue
		}

		//处理用户连接: 
		// 1. 新连接加入 map
		// 2. 开启等待读取 Client.C 写回客户端的协程，等待读取Client.C
		// 3. 广播该客户端上线了
		go HandleConn(conn)

	}
}

//死循环，阻塞等待，写消息(写入 Client.C)
func sendMessage2ClientChan() {
	for {
		msg := <- message //阻塞等待读取消息

		for _, cli := range onlineMap {
			cli.C <- msg //写
		}
	}
}

//读取 Client.C 转发给具体的 CLient
func writeMsg2Client(cli Client, conn net.Conn) {
	for msg := range cli.C { //读
		conn.Write([]byte(msg + "\n"))
	}
}


//处理新连接
func HandleConn(conn net.Conn) { //处理用户连接
	defer conn.Close()  //处理完就关闭连接

	//获取客户端信息，加入 map
	cliAddr := conn.RemoteAddr().String()  //作为key
	////默认情况 Client.Name 就是地址名字
	cli := Client{make(chan string), cliAddr, cliAddr}

	//把结构体添加到 map
	onlineMap[cliAddr] = cli

	 //等待读取 Client.C，转发给 Client 对端
	go writeMsg2Client(cli, conn)

	//写 mesage 给具体的 cli.C
	message <- makeMsg(cli, "上线了。")
	//写给自己端，提示我的用户是谁(只给自己写) ---- 后面的改名功能
	cli.C <- makeMsg(cli, cli.Name + "初来乍到，请多指教")

	//检测是否主动退出
	isQuit := make(chan bool)
	isAlive := make(chan bool)

	//新开协程接收用户Client的数据
	go func() {
		buf := make([]byte, 1024*2)
		for {
			n, err := conn.Read(buf)
			if n==0 { //对端断开或者其他连接问题
				log.Println("conn.Read err = ", err) //某个用户断开
				isQuit <- true
				return //该协程结束，而不是server结束
			}
			//转发读到的消息
			msg := string(buf[:n-1]) //读多少，转发多少
			len := len(msg)

			if len == 3 && msg == "who" {
				//遍历 map，给当前用户发送所有成员
				conn.Write([]byte("-----------\n用户列表如下: \n"))
				for _, tmp := range onlineMap {
					usrStr := tmp.Addr + ":" + tmp.Name +"\n"
					conn.Write([]byte(usrStr))
				}
				conn.Write([]byte("-----------\n"))

			} else if len >=8 && msg[:6] == "rename" { //rename|骨傲天
				newName := strings.Split(msg, "|")[1]
				cli.Name = newName
				onlineMap[cliAddr] = cli //因为map存储的是副本
				conn.Write([]byte("rename 完毕。\n"))
			} else {
				message <- makeMsg(cli, msg)
			}

			isAlive <- true //流程走到此，说明前面 conn.Read() 阻塞读取的内容有数据
		}
	}()

	for {
		select {
		case <- isQuit:
			//移除当前用户
			delete(onlineMap, cliAddr) //从 map 移除
			message <- makeMsg(cli, "下线了。")
			return
		case <- isAlive:
			//do nothing
		case <- time.After(10 * time.Second): //超时了 (其他分支都没有动作)
			delete(onlineMap, cliAddr) //从 map 移除
			message <- makeMsg(cli, "被踢下线，理由：潜水暗中观察。")
			return
		}

	}//连接还不能关闭

}

func makeMsg(cli Client, str string) (buf string) {
	buf = "[ " + cli.Addr + " ] " + cli.Name + " : " + str
	return
}



