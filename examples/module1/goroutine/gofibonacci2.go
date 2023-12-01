package main

import (
	"fmt"
)

func fibonacci1(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func mainf2() {
	// 带缓冲区的通道 不会阻塞 除非空间耗尽
	// chan 都通过make 来创建 引用类型
	c := make(chan int, 10)
	go fibonacci1(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
