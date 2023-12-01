package main

import (
	"fmt"
	"os"
)

func tel1(ch chan int, quit chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	//循环结束写入值
	quit <- true
}

// 使用一个额外的通道传递给协程，然后在结束的时候随便放点什么进去。
// main() 线程检查是否有数据发送给了这个通道，如果有就停止：goroutine_select.go
/*
// 往通道发送数据的时候
***********在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！如果没有 case，select 就会一直阻塞。

select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出。

*/
func mains2() {
	//make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel
	ch := make(chan int)
	quit := make(chan bool)

	go tel1(ch, quit)
	for {
		//
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit: //当发送者可用的时候 也就是tel1方法执行完成 这个是在最后 所以根 close 作用相同
			os.Exit(0)
		}
	}
}
