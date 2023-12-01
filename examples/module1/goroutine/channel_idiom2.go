package main

import (
	"fmt"
	"time"
)

func mainidiom2() {
	suck3(pump3())
	//防止释放
	time.Sleep(1e9)
}

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

/*
它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。
很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），
而且必须在写入完成后才关闭。suck 函数可以这样写，且在协程中调用这个动作

*/
func suck3(ch chan int) {
	//另外一个协程必须写入 ch（不然代码就阻塞在 for (这里指pubmp3函数)循环了），而且必须在写入完成后才关闭
	go func() {
		//这个 跟切片循环不一样 这里必须 这里只有 value 没有index
		// 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
