// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
)

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

// 函数类型
type binOp func(a, b int) int

func run(op binOp, req *Request) {
	//往通道写入值 eq.replyc 不取值的话会一直卡住
	req.replyc <- op(req.a, req.b)
	fmt.Println("------------")
}

func server(op binOp, service chan *Request) {
	for {
		// 没有可用发送者 在当前线程阻塞
		req := <-service // requests arrive here
		// start goroutine for request:
		go run(op, req) // don't wait for op
	}
}

func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request) //make 创建通道
	// 开一个 goroutine 来处理
	go server(op, reqChan)
	return reqChan
}

/*这个程序只开启 100 个 Goroutines 。执行 100000 个 Goroutines 的程序，
甚至可以看到它在几秒钟内完成。这说明了 Goroutines 是有多么的轻量：如果我们启动相同数量的实际线程，程序将很快崩溃。
;*/
func main() {
	// 我们发送 100 个请求，并在所有请求发送完毕后，再逐个检查其返回的结果：
	adder := startServer(func(a, b int) int { return a + b })
	const N = 100
	// 当声明数组时所有的元素都会被自动初始化为默认值 0。
	var reqs [N]Request //{0 0 <nil>} 默认值
	fmt.Println("-----", reqs)
	for i := 0; i < N; i++ {
		req := &reqs[i] //指针
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		//100次
		// 这里如果没有可用的接受者会阻塞 调用第一次就卡住等待可用接受者
		adder <- req
	}
	// checks:
	// for i := N - 1; i >= 0; i-- { // doesn't matter what order
	// 	if <-reqs[i].replyc != N+2*i {
	// 		fmt.Println("fail at", i)
	// 	} else {
	// 		fmt.Println("Request ", i, " is ok!")
	// 	}
	// }
	for _, v := range reqs {
		fmt.Println(<-v.replyc)
	}
	fmt.Println("done")
}
