package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	// fmt.Printf format %d has arg s of wrong type string go vet main.go
	// fmt.Printf("Here is the string s: %d\n", s) // prints same string
	fmt.Printf("Here is the string s: %s\n", s)

	const i = 5
	// \你不能得到一个文字或常量的地址
	// ptr := &i //error: cannot take the address of i
	// ptr2 := &10 //error: cannot take the address of 10
	// 在大多数情况下 Go 语言可以使程序员轻松创建指针，并且隐藏间接引用，如：自动反向引用。
	// 对一个空指针的反向引用是不合法的，并且会使程序崩溃：
	var p1 *int = nil
	*p1 = 0
}
