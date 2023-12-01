package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本A:
	// 下标
	for ix := range values { // ix是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
	}
	// 0 1 2 3 4
	fmt.Println()
	// 版本B: 和A版本类似，但是通过调用闭包作为一个协程
	for ix := range values {
		//
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	// go异步执行 因为协程可能在循环结束后还没有开始执行，而此时 ix 值是 4。
	fmt.Println() //4 4 4 4 4
	time.Sleep(5e9)
	// 版本C: 正确的处理方式
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	// 4 2 0 1 3 根据协程被执行顺序 每次都不一样
	fmt.Println()
	time.Sleep(5e9)
	// 版本D: 输出值:
	for ix := range values {
		// 因为版本 D 中的变量声明是在循环体内部，所以在每次循环时
		// ，这些变量相互之间是不共享的，所以这些变量可以单独的被每个闭包使用。
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	// 14 12 10 11 13
	time.Sleep(1e9)
}
