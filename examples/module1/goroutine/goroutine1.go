package main

import (
	"fmt"
	"time"
)

/*
Go 协程（goroutines）和协程（coroutines）goroutine 也是协程比其他的更强大。。
----------------------------------------------------------------
Go 协程 比其他更强大
Go 协程意味着并发（或者可以以并行的方式部署），协程一般来说不是这样的
Go 协程通过通道来通信；协程通过让出和恢复操作来通信
Go 协程比协程更强大，也很容易从协程的逻辑复用到 Go 协程。
*****并行是一种通过使用多处理器以提高速度的能力。所以并发程序可以是并行的，也可以不是。
GOMAXPROCS 等同于（并发的）线程数量，在一台核心数多于 1 个的机器上，会尽可能有等同于核心数的线程在并行运行。
*/
func main1() {
	/*我们让 main() 函数暂停 10 秒从而确定它会在另外两个协程之后结束。
	如果不这样（如果我们让 main() 函数停止 4 秒），main() 会提前结束，
	longWait() 则无法完成。如果我们不在 main() 中等待，
	协程会随着程序的结束而消亡。

	*/
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")

}

func longWait() {
	fmt.Println("Beginning longWait()")
	//五秒
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
