package main

import (
	"fmt"
	"time"
)

func mainblock2() {
	ch1 := make(chan int)
	go pump1(ch1)
	go suck(ch1)
	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
