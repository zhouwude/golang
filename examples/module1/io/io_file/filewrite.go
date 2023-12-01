package main

import (
	"fmt"
	"os"
)

func mainwrite() {
	//0 stdin 标准输入 1 stdout 标准输出 2stderr 错误输出
	os.Stdout.WriteString("hello, world\n")
	// O_CREATE没有会创建
	f, _ := os.OpenFile("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/test.dat", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	_, error := f.WriteString("hello, world in a file\n")
	if error != nil {
		fmt.Println(error.Error())
	}
}
