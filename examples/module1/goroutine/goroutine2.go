package main

import (
	"fmt"
	"time"
)

// channel是引用类型，定义的时候 加上 chan string， chan int 可以是任何类型包括interface{}类型
// 引用类型未初始化默认 nil
func main() {
	ch := make(chan string)
	// 如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。
	go sendData(ch)
	go getData(ch)
	// 协程会随着程序的结束而消亡 防止协程释放掉
	time.Sleep(1e9) //
}

func sendData(ch chan string) {
	// 流向通道（发送） 用通道 ch 发送变量前面是通道变量
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string //没定义 :=
	// 接收值 chan在操作符后面
	// time.Sleep(2e9)

	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}
