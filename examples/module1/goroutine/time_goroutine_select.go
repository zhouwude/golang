package main

import (
	"fmt"
	"time"
)

func maintime() {
	// 当你想返回一个通道而不必关闭它的时候这个函数非常有用：它以 d 为周期给返回的通道发送时间，d 是纳秒数。
	tick := time.Tick(1e8)
	// 在 Duration d 之后，当前时间被发到返回的通道
	boom := time.After(5e8)
	// select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break return os.Exit(int) 语句使循环退出。
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			// return 管不循环
			return
			// 在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！
			// 如果没有 case，select 就会一直阻塞。
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
