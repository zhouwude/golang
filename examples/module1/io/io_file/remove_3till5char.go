package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func mainq() {
	inputFile, _ := os.Open("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/goprogram")
	outputFile, _ := os.OpenFile("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/goprogramT", os.O_WRONLY|os.O_CREATE, 0666) //没有会创建
	defer inputFile.Close()
	defer outputFile.Close()
	//读写缓冲区 实现了 Writer Reader interface

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	var outputString string
	for {
		// inputString, readerError := inputReader.ReadString('\n')
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		//fmt.Printf("The input was: --%s--", inputString)
		if len(inputString) < 3 {
			outputString = "\r\n"
		} else if len(inputString) < 5 {
			outputString = string([]byte(inputString)[2:len(inputString)]) + "\r\n"
		} else {
			// string([]byte) 字节数组直接转化成 string
			outputString = string([]byte(inputString)[2:5]) + "\r\n"
		}
		//fmt.Printf("The output was: --%s--", outputString)
		_, err := outputWriter.WriteString(outputString)
		//fmt.Printf("Number of bytes written %d\n", n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	outputWriter.Flush()
	fmt.Println("Conversion done")
}
