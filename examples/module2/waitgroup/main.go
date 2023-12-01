package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 等待一组 goroutine 返回
	// waitBySleep()
	// waitByChannel()
	waitByWG()
}

//通过Sleep等待
func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

// 通过通道等待
func waitByChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		// 会拷贝每一个 i 的值 异步执行
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		//读取值会阻塞 没有可用接受者 即使是带缓冲通道
		<-c
	}
}

// 通过Waitgroup等待
func waitByWG() {
	wg := sync.WaitGroup{}
	fmt.Println(wg) //{{} 0 0} 解构体初始化不需要一定给变量赋值不赋值默认给类型的零值
	wg.Add(100)     //限制个数
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			// 退出协程 类似iOS dispatch_group
			wg.Done()
		}(i)
	}
	// 阻塞
	wg.Wait()
}
