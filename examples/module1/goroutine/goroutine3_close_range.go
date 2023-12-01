package main

import "fmt"

func main3() {
	ch := make(chan string)
	go sendData1(ch)
	getData1(ch)
}

// 第一个可以通过函数 close(ch) 来完成：这个将通道标记为无法通过发送操作
// <- 接受更多的值；给已经关闭的通道发送或者再次关闭都会导致运行时的 panic
func sendData1(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	//通道无需每次关闭
	//关闭通道是为了告诉接受者通道再无新数据发送
	//只有发送方需要关闭通道
	close(ch) //
}

func getData1(ch chan string) {
	// 循环 防止程序结束
	for {
		// 如何来检测通道收到没有被阻塞（或者通道没有被关闭）？
		// input 的值
		// *******给已经关闭的通道发送或者再次关闭都会导致运行时的 panic 没关闭的话就会阻塞而不是抛出错误。
		input, open := <-ch
		if !open {
			fmt.Println("closed")
			break
		}
		fmt.Printf("%s ", input)
		// Washington Tripoli London Beijing Tokio closed
	}
	// 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：只有 value 没用 subscript

	// for input := range ch {
	// 	fmt.Printf("%s ", input)
	// }
}
