package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// go可以类型推断 不需要直接声明类型
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]byte, 4096) //切片
		read          = true
		count         = 0
	)
	// 创建一个socket
	con, err := net.Dial("tcp", remote)
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}
	// 发送我们的消息，一个http GET请求
	// io.WriteString函数的作用是将字符串写入到指定的io.Writer对象中。
	// 它将字符串作为字节序列进行处理，并将其逐个写入到io.Writer中。这个函数在底层使用了io.Writer接口的Write方法来实现字符串的写入操作。

	// 使用io.WriteString函数可以方便地将字符串写入到各种输出设备或数据流中，比如文件、网络连接、内存缓冲区等。
	// 只要这些对象实现了io.Writer接口，就可以使用io.WriteString函数来进行写入操作。
	// 往 conn对象中 写入数据
	io.WriteString(con, msg)
	// 读取服务器的响应
	for read {
		count, err = con.Read(data)
		read = (err == nil) //判断条件
		//复制一个切片
		fmt.Println(string(data[0:count]))
	}
	con.Close()
}
