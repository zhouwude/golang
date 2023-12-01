package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

/*服务器运行的时候，你无法编译 / 连接同一个目录下的源码来产生一个新的版本，
因为 server.exe 正在被操作系统使用而无法被替换成新的版本。

下边这个版本的 simple_tcp_server.go 从很多方面优化了第一个 tcp
服务器的示例 server.go 并且拥有更好的结构，它只用了 80 行代码！

*/
const maxRead = 25

func main() {
	flag.Parse() //第一个参数可用 不再是文件名
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	// 服务器地址和端口不再是硬编码，而是通过命令行传入参数并通过 flag 包来读取这些参数。
	// 这里使用了 flag.NArg() 检查是否按照期望传入了 2 个参数：
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}
func initServer(hostAndPort string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}
func connectionHandler(conn net.Conn) {
	// conn.RemoteAddr() 获取到客户端的地址
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	// 发送改进的 go-message 给客户端
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		// ibuf[0:maxRead]属于重组切片 内存和ibuf 一样
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // 防止溢流
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
			// syscall 是低阶外部包，用来提供系统基本调用的原始接口。
			// 它们返回整数的错误码；类型 syscall.Errno 实现了 Error 接口。
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}
func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+string(rune(wrote))+" bytes.")
}
func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 { //最后 跳出循环
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}
func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}
