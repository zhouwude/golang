package main

import (
	"fmt"
	"time"
)

// 简单的超时模板
func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int)
	// <-timeout
	go func() {
		time.Sleep(1e9) // one second
		timeout <- true
	}()
	select {
	case <-ch:
	// a read from ch has occurred
	case <-timeout:
		fmt.Println("time out") //time out
		// the read from ch has timed out
	}
}
