package main

import (
	"fmt"
	"time"
)

func mainblock3() {
	//不带缓冲的通道
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}() //立即执行
	fmt.Println("sending", 10)
	//没有可用的接受者 阻塞了 知道 15s之后出现接收者阻塞才恢复
	c <- 10
	fmt.Println("sent", 10)
	/*
				sending 10
				received 10
		sent 10
	*/
}
