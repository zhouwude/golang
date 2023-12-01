package main

import "fmt"

var done = make(chan bool)
var msgs = make(chan int)

func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	//相当于一个开关的通道 阻塞
	done <- true
}

func consume() {
	for {
		msg := <-msgs
		fmt.Print(msg, " ")
	}
}

func mainconsumer2() {
	go produce()
	go consume()
	//阻塞防止 结束
	<-done
}
