package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
	// if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
	// 	err = fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
	// 	return
	// }
	fmt.Println()
	arr := [100]byte{}
	//100 100
	// 切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量。
	fmt.Println(len(arr), cap(arr))
	s := arr[:10]
	fmt.Println(s, len(s), cap(s)) //[0 0 0 0 0 0 0 0 0 0] 10 100
	s = s[:len(s)+1]
	fmt.Println(s, len(s), cap(s)) //[0 0 0 0 0 0 0 0 0 0 0] 11 100 扩容

	// **创建一个 error
	fmt.Println(errors.New("this is error"))
	fmt.Println(fmt.Errorf("%s", "this is a error"))
	// this is error
	// this is a error
}

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, errors.New("math - square root of negative number")
// 	}
// 	// implementation of Sqrt
// }
