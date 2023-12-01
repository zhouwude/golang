package main

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/module1/init/a"
	_ "github.com/cncamp/golang/examples/module1/init/b"
)

/*
//这里包含 a 包 和b包 先导入
按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

init方法只会被执行一次 不管被依赖多次 ，线程安全
*/
func init() {
	fmt.Println("main init")
}
func main() {
	/*
	   init from b    b模块最先执行
	   init from a
	   main init
	*/
}
