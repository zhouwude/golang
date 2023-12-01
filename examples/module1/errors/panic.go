package main

import (
	"fmt"
	"os"
)

func maincopy() {
	// 当错误条件（我们所测试的代码）很严苛且不可恢复，程序不能继续运行时，可以使用 panic 函数产生一个中止程序的运行时错误
	// fmt.Println("Starting the program")
	// panic("A severe error occurred: stopping the program!")
	// fmt.Println("Ending the program")
	check()
}
func check() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println(3)
	}()
	// 逆序执行---
}
