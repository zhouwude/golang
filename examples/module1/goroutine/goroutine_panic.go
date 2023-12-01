package main

import (
	"fmt"
	"time"
)

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	// close(ch) 如果发生这种情况:panic:所有例程都处于休眠状态-死锁!
}

func mainpanic() {
	var ok = true
	ch := make(chan int)

	go tel(ch)
	time.Sleep(1e9)
	// 一直循环
	for ok {
		// 判断关闭状态 通道没有被关闭
		i, ok := <-ch //阻塞了 会panic
		if ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		} else {
			fmt.Println("阻塞死锁")
		}
	}
}
