// rpc_server.go
// after client-exits the server shows the message:
//       1:1234: The specified network name is no longer available.
//       2011/08/01 16:19:04 rpc: rpc: server cannot decode request: WSARecv tcp 127.0.0.
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/cncamp/golang/examples/module1/net/rpc_objects/rpc_objects"
)

// 它提供了通过网络连接进行函数调用的便捷方法。只有程序运行在不同的机器上它才有用。
func mains() {
	calc := new(rpc_objects.Args)
	// 服务器创建一个用于计算的对象，并且将它通过 rpc.Register(object) 注册，
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	/*对每一个进入到 listener 的请求，都是由协程去启动一个 http.Serve(listener, nil)
	，为每一个传入的 HTTP 连接创建一个新的服务线程。
	我们必须保证在一个特定的时间内服务器是唤醒状态
	*/
	go http.Serve(listener, nil)
	time.Sleep(1000e9) //
}

/* Output:
Starting Process E:/Go/GoBoek/code_examples/chapter_14/rpc_server.exe ...

** after 5 s: **
End Process exit status 0
*/
