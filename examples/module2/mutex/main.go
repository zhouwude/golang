package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go rLock()
	go wLock()
	go lock()
	time.Sleep(5 * time.Second)
}

func lock() {
	lock := sync.Mutex{}
	// 当有多个 defer 行为被注册时，它们会以逆序执行 （类似栈，即后进先出）：
	// 这里在循环中 定义 defer 相当于定义了多个 defer方法 有多少个循环定义了多少个方法
	// 但是在循环结尾处的 defer 没有执行，所以文件一直没有关闭 函数返回的时候才执行

	// defer 仅在函数返回时才会执行，在循环的结尾或其他一些有限范围的代码内不会执行。
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock:", i)
	}
}

func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rLock:", i)
	}
}

func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wLock:", i)
	}
}
