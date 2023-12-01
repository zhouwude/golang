// degob.go
package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address2 struct {
	Type    string
	City    string
	Country string
}

type VCard2 struct {
	FirstName  string
	LastName   string
	Addresses2 []*Address2
	Remark     string
}

// var content1 string
var vc VCard2

func maindegob() {
	// using a decoder:
	file, _ := os.Open("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc) //{  [] }
}
