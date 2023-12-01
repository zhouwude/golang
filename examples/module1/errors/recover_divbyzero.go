package main

import (
	"fmt"
)

func badCall1() {
	// 这里不是用户自定义的 panic 是内部定义的 panic
	a, b := 10, 0
	n := a / b //取整数 b为0会崩溃
	fmt.Println(n)
}

func test1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}

	}()
	badCall1()
	fmt.Printf("After bad call\r\n")
}

func maindiv() {
	fmt.Printf("Calling test\r\n")
	test1()
	fmt.Printf("Test completed\r\n")
	/*Calling test
	Panicing runtime error: integer divide by zero
	Test completed*/
}
