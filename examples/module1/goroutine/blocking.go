package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

// 互相阻塞 形成死锁
func mainblocking() {

	out := make(chan int)
	// 在 main函数线程 创建无缓冲的通道 没有可用接受者直接 阻塞了当前线程go f1(out)无法执行
	out <- 2
	go f1(out) //阻塞了 走不到这里了
}
