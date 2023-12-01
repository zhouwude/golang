// hash_md5.go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func mainmd5() {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "test")
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}
