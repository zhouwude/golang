package main

import "fmt"

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		// 在 select 中使用发送操作并且有 default 可以确保发送不被阻塞！如果没有 case，select 就会一直阻塞。
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit: //不需要接收值直接 -<就行
			fmt.Println("quit")
			return
		}
	}
}

func mainfs() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
