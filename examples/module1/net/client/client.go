package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
如果在服务器没有开始监听的情况下运行客户端程序，
客户端会停止并打印出以下错误信息：对tcp 127.0.0.1:50000发起连接时产生错误：
由于目标计算机的积极拒绝而无法创建连接
*/
func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}
	//bufio带缓冲区的文件读取
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	// 读到换行符
	clientName, _ := inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s", clientName)
	trimmedClient := strings.Trim(clientName, "\n") // Windows 平台下用 "\r\n"，Linux平台下使用 "\n"
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--s%--", input)
		// fmt.Printf("trimmedInput:--s%--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		if _, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput)); err != nil {
			fmt.Println(err.Error())
		}
	}
}
