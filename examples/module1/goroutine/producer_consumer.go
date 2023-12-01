// goroutines2.go
package main

import "fmt"

// integer producer:
// chan<- int是一个只发送通道，只能用于发送数据到通道中。
// <- chan int只能接收值
// chan int是一个双向通道，既可以用于发送数据到通道中，也可以从通道中接收数据。
// 通道的方向
// 通过使用方向注解来限制协程对通道的操作
func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	//关闭通道因为关闭通道是发送者用来表示不再给通道发送值了
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	//for range循环 channel 变量
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	// 生产者使 main函数结束 读完数据
	done <- true
}

func maincon() {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)
	//不带缓冲 done 通道变量没有可用发送者则阻塞
	// 等待两个协程完成后再结束。语句作用
	<-done //主函数 这里没有生产者 所以阻塞 main函数不会释放
}
