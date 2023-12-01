package main

import (
	"fmt"
	"time"
)

/*
select 做的就是：选择处理列出的多个通信情况中的一个。

如果都阻塞了，会等待直到其中一个可以处理 一般放在for 中
如果多个可以处理，随机选择一个
如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。

*/
func mains4() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump4(ch1)
	go pump5(ch2)
	// go suck4(ch1, ch2)
	// 不调用上面那句话循环都被阻塞
	/*ch - for
	ch1 - for*/

	time.Sleep(1e9)
}

func pump4(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("ch - for")
		// 没有接收者 就阻塞知道接受者可用
		ch <- i * 2
	}
}

func pump5(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("ch1 - for")
		ch <- i + 5
	}
}

func suck4(ch1, ch2 chan int) {
	// 读取数据释放循环
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
