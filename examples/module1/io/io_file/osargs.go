package main

import (
	"fmt"
	"os"
	"strings"
)

func maina() {
	who := "Alice "
	if len(os.Args) > 1 {
		fmt.Println(os.Args[0]) //第一个应该是路径 os.Args[0] 放的是程序本身的名字
		who += strings.Join(os.Args[1:], " ")
	}
	// /var/folders/57/r34ywnhn7l7bt5kj19qkwbw00000gn/T/go-build3013585383/b001/exe/osargs
	// Good Morning Alice 1 2 3
	fmt.Println("Good Morning", who)
}
