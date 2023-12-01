package main

import (
	"fmt"
	"os"
	"path/filepath"
	//	"io/ioutil"
	//	"strings"
)

/*如果数据是按列排列并用空格分隔的，你可以使用 fmt 包提供的以 FScan 开头的一系列函数来读取他们。
请看以下程序，我们将 3 列的数据分别读入变量 v1、v2 和 v3 内，然后分别把他们添加到切片的尾部。
;*/
func main3() {
	file, err := os.Open("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/product2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
	filename := filepath.Base("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/product2.txt")
	fmt.Println(filename) //product2.txt
}
