package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main1() {
	// path, _ := os.Getwd()

	// fmt.Println("Path:", path) //查找路径 /Users/zhouwude/Desktop/golang-master
	inputFile, inputError := os.Open("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/input.txt")

	// path := "./input.txt"
	// path, _ = filepath.Abs(path)
	// fmt.Println("Path:", path)
	// inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	// 在 return 变量赋值之后 return之前执行
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		// inputReader.ReadLine()

		fmt.Printf("The input was: %s", inputString)
		// io.EOF 文件结束
		if readerError == io.EOF {
			return
		}
	}

}
