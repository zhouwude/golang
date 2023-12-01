package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	for i := 0; ; i++ {
		// 定义一个死循环
		fmt.Println(i)
	}
	fullString := "hello world"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}
	s := []byte{1, 2, 3, 4}
	s1 := s[:2]

	fmt.Println(s1)
	fmt.Println(append(s1, 100))

	fmt.Println(s)               //[1 2 100 4]
	fmt.Printf("%p---%p", s, s1) //0xc0000b004c---0xc0000b004c
	fmt.Println(cap(s), cap(s1)) //44
}
