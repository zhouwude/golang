package main

import (
	"bufio"
	"fmt"
	"os"
)

/*可以看到，OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符 “|” 连接），使用的文件权限。

我们通常会用到以下标志：

os.O_RDONLY：只读
os.O_WRONLY：只写
os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0。
在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用 0。而在写文件时，
不管是 Unix 还是 Windows，都需要使用 0666。
0666八进制

*/
func main6() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string
	// 0666 110 110 110有读写权限
	outputFile, outputError := os.OpenFile("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	// 在缓冲写入的最后千万不要忘了使用 Flush()，否则最后的输出不会被写入。
	outputWriter.Flush()
}
