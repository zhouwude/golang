package main

import (
	"fmt"
)

/*

在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
如果存在 init 函数的话，(*****init方法会在 main 函数之前被执行*****)则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
如果当前包是 main 包，则定义 main 函数。
然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。


Go 程序的执行（程序启动）顺序如下：

按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

*/
const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
}
