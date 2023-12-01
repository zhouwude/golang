// gob2.go
package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address1 struct {
	Type    string
	City    string
	Country string
}

type VCard1 struct {
	FirstName  string
	LastName   string
	Addresses1 []*Address1
	Remark     string
}

var content string

// 这会产生一个文本可读数据和二进制数据的混合，当你试着在文本编辑中打开的时候会看到。
func maingob2() {
	pa := &Address1{"private", "Aartselaar", "Belgium"}
	wa := &Address1{"work", "Boom", "Belgium"}
	vc := VCard1{"Jan", "Kersschot", []*Address1{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("/Users/zhouwude/Desktop/golang-master/examples/module1/io/io_file/vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
