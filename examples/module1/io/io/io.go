package main

import (
	"bufio"
	"fmt"
	"os"

	wordlettercount "github.com/cncamp/golang/examples/module1/io/word_letter_count"
)

// io 包里的 Readers 和 Writers 都是不带缓冲的，
// bufio 包里提供了对应的带缓冲的操作，在读写 UTF-8 编码的文本文件时它们尤其有用
// 因式分解的方式
// var (
// 	firstName, lastName, s string
// 	i                      int
// 	f                      float32
// 	input                  = "56.12 / 5212 / Go"
// 	format                 = "%f / %d / %s"
// )
var inputReader *bufio.Reader
var input string
var err error

func main() {
	// fmt.Println("Please enter your full name: ")
	// fmt.Scanln(&firstName, &lastName)
	// // fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n') //遇到换行符结束

	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
	fmt.Printf("Your name is %s", input)
	//
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip!")
	case "Chris\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\n":
		fallthrough
	case "Ivo\n":
		fallthrough
	case "Chris\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\n", "Ivo\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
	wordlettercount.Start()

}
