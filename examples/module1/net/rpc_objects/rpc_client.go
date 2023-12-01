package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/cncamp/golang/examples/module1/net/rpc_objects/rpc_objects"
)

const serverAddress = "localhost"

func main() {
	// 去创建连接的客户端
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpc_objects.Args{7, 8}
	var reply int
	// 当客户端被创建时，它可以通过
	/*client.Call("Type. Method", args, &reply) 去调用远程的方法，其中 Type
	与 Method 是调用的远程服务器端被定义的类型和方法， args 是一个类型的初始化对象，
	reply 是一个变量，使用前必须要先声明它，它用来存储调用方法的返回结果。

	*/

	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
