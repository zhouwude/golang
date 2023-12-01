// lazy_evaluation.go
package main

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	//不停的发送数据由于没有接受这故阻塞了
	go func() {
		for {
			fmt.Println("--for")
			yield <- count
			fmt.Println("--for end")
			count++
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func mainlz() {
	resume = integers()
	fmt.Println("--gen")
	fmt.Println(generateInteger()) //=> 0 第一个接受者 yield继续执行知道发现第二个 consumer
	// fmt.Println(generateInteger()) //=> 1
	// fmt.Println(generateInteger()) //=> 2
	/*--gen
	--for
	--for end
	--for
	0*/
}
