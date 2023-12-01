package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")

	// 创建 listener
	// 127.0.0.1 port 50000
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	fmt.Printf("Listen %s ...", listener.Addr().String())
	// 监听并接受来自客户端的连接
	// go 会为每一个客户端产生一个协程用来处理请求。我们需要使用 net 包中网络通信的功能。
	// 它包含了用于 TCP/IP 以及 UDP 协议、域名解析等方法。
	for {
		// 用一个无限 for 循环的 listener.Accept() 来等待客户端的请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		// 异步处理每一个客户端发来的消息
		go doServerStuff(conn)
	}
}

/*独立的协程使用这个连接执行 doServerStuff()，
开始使用一个 512 字节的缓冲 data 来读取客户端发送来的数据并且把它们打印到服务器的终端，
len 获取客户端发送的数据字节数；当客户端发送的所有数据都被读取完成时，协程就结束了。
这段程序会为每一个客户端连接创建一个独立的协程。必须先运行服务器代码，再运行客户端代码。

*/
func doServerStuff(conn net.Conn) {

	for {
		buf := make([]byte, 512) //512字节来读取客户端的内容
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Println()
		//byte 切片 转 string
		fmt.Printf("Received data: %v", string(buf[:len]))
		// sort.Reverse(new(V))

	}
}

// Less(i, j int) bool

// 	// Swap swaps the elements with indexes i and j.
// 	Swap(i, j int)
type V struct {
}

func (v *V) Len() (l int) {
	l = 3
	return
}
func (v *V) Less(i, j int) bool {
	return i < j
}
func (v *V) Swap(i, j int) {

}
