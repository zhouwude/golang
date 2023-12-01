// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	// 没有接受者 是阻塞的
	for i := 2; ; i++ {
		//无缓冲区的通道没有空间来存储任何内容,必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者
		ch <- 2 // Send 'i' to channel 'ch'.
	}
}

//从通道'in'复制值到通道'out'，
//移除那些能被'素数'整除的数。
func filter(in, out chan int, prime int) {
	for {
		i := <-in         // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 { //奇数
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

/*
// chan<- int是一个只发送通道，只能用于发送数据到通道中。
// <- chan int只能接收值
// chan int是一个双向通道，既可以用于发送数据到通道中，也可以从通道中接收数据。
// 通道的方向
// 通过使用方向注解来限制协程对通道的操作
*/
// The prime sieve: Daisy-chain filter processes together.
func mainsv1() {
	//
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch // ch 变成了 ch1接收ch1的值
		fmt.Print(prime, " ")
		ch1 := make(chan int) //创建 ch1
		go filter(ch, ch1, prime)
		// // 记录上一次的通道
		ch = ch1
	}
}
