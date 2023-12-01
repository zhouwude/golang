package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

// tt := 100 简洁格式只能在函数体内
var tt = 100 //只能这么写
// defer 仅在函数返回时才会执行，在循环的结尾或其他一些有限范围的代码内不会执行。
/*

for _, file := range files {
    if f, err = os.Open(file); err != nil {
        return
    }
    // 对文件进行操作
    f.Process(data)
    // 关闭文件
    f.Close()
 }


*/
func main() {

	// 逆序
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// loopFunc()
	// time.Sleep(time.Second)
	const a int = 10 //在 Go 语言中，你可以省略类型说明符 [type]，因为编译器可以根据常量的值来推断其类型。
	fmt.Println(a)
	var b = 10
	fmt.Println(&b) //0xc0000b0010
	fmt.Printf("%v", &b)
	// 枚举
	const (
		a1 = iota
		b1 = iota
		c1 = iota
	)

	const (
		a11 = iota
		b11
		c11
	)
	//  ioto 是零值
	const (
		a111 = iota
		b111 = 10
		c111
	)
	// 同时声明多个变量 const 也可以
	// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，
	// float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。记住，所有的内存在 Go 中都是经过初始化的。

	// var (
	// 	x int
	// 	y bool
	// 	z string
	// )
	fmt.Println(os.Getenv("GOROOT")) ///Users/zhouwude/.gvm/gos/go1.18.3
	fmt.Println(runtime.GOOS)        //darwin
	// %p 打印变量的内存地址--------
	x, y := 10, 5
	x, y = y, x
	fmt.Println(x, y)
	// 直接交换变量的值

}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// go func(i int) {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("loopFunc:", i)
		// }(i)
	}
}
