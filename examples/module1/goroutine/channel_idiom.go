// 在 channel_block2.go 加入工厂这种模式
package main

import (
	"fmt"
	"time"
)

func mainidiom1() {
	stream := pump2()
	go suck1(stream)
	time.Sleep(1e9)
}

// 通道的工厂模板：以下函数是一个通道工厂，启动一个匿名函数作为协程以生产通道：
func pump2() chan int {
	ch := make(chan int)
	// 匿名立即执行函数的协程
	go func() {
		// 传入值
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck1(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
