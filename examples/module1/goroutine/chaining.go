package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

// last 相当于上一个的意思
func f(last, current chan int) {
	temp := 1 + <-current // 当 current 没有值时,这里会被阻塞
	// 这里的 last 相当与 current的上一个通道调用上一层级的 f()函数 一直到最上层。
	last <- temp
}
func mainchan() {
	first := make(chan int) //引用类型
	fmt.Println(*ngoroutine)
	last := first //拷贝一个 跟first地址不一样
	// first 0xc0000aa018------last 0xc0000aa028100000
	fmt.Printf("first %v------last %v", &first, &last)
	// 执行完 才会执行 循环一万次
	for i := 0; i < *ngoroutine; i++ {
		current := make(chan int)
		// 第一次传入的其实是 first 一样的 channel这里复制一个是防止后面被重新赋值
		// 将上一次循环创建的 chan,和本次循环的 chan 一起交给函数, 函数会帮我们完成 last <- 1+ <- curr 的过程
		go f(last, current) //不在 main线程 每一个函数调用都保存在栈上 而且被阻阻塞

		// 记录本次循环中的 right,给下一次循环创建使用
		last = current
	}

	// 开始链接
	last <- 0 // 发送值 看看有没有可用的接收者 last是最后一个通道 发送值 是的阻塞函数执行
	//这里的 first 是循环第一次执行的左边的通道。计算累计值。。。
	x := <-first // wait for completion 等待完成

	fmt.Println(x)
	// 结果： 100000 ， 大约 1,5s （我实际测试只用了不到200ms）
}
