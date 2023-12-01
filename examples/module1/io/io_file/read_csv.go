package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func maincvs() {
	//os.FileMode
	fmt.Println("ModeDir value: ", os.ModeDir)        //d---------是个目录
	fmt.Println("ModeType value: ", os.ModeType)      //dLDpSc?---------
	fmt.Println("FileMode value: ", os.ModePerm)      //-rwxrwxrwx
	fmt.Println(strconv.ParseInt("101110101", 2, 64)) //373
	fmt.Println(strconv.ParseInt("4e", 16, 64))       //78
	// 格式化

	fmt.Println(strconv.FormatInt(0x445e, 2)) //100010001011110
	bks := make([]*Book, 1)
	file, err := os.Open("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/product.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// read one line from the file:
		line, err := reader.ReadString('\n')
		readErr := err
		// remove \r and \n so 2(in Windows, in Linux only \n, so 1):
		line = line[:len(line)-2]
		//fmt.Printf("The input was: -%s-", line)

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//fmt.Printf("The quan was:-%s-", strSl[2])
		book.quantity, err = strconv.Atoi(strSl[2])
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}

		bks = append(bks, book)

		if readErr == io.EOF {
			break
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
