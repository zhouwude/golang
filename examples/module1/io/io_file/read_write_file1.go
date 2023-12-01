package main

import (
	"fmt"
	"io/ioutil"
)

// 将整个文件的内容读到一个字符串里：
func main2() {
	inputFile := "/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/product.txt"
	outputFile := "/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/product_copy.txt"
	/*
			*****将****整个文件****的内容读到一个字符串里：

		如果您想这么做，可以使用 io/ioutil 包里的 ioutil.ReadFile() 方法，
		该方法第一个返回值的类型是 []byte，里面存放读取到的内容，
		第二个返回值是错误，如果没有错误发生，第二个返回值为 nil。请看示例 12.5。
		类似的，函数 WriteFile() 可以将 []byte 的值写入文件。


	*/
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		// 抛出错误 -----------
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))

	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
