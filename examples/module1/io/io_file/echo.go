package main

import (
	"flag" // 命令行选项解析器
	"os"
)

// 定义了一个默认值是 false 的 flag
// 当在命令行出现了第一个参数（这里是 “n”），flag 被设置成 true
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
const (
	Space   = " "
	Newline = "\n"
)

func mainE() {
	flag.PrintDefaults() //打印 flag 的使用帮助信息 -n	print newline
	// Parse() 之后 flag.Arg(i) 全部可用，flag.Arg(0) 就是第一个真实的 flag，而不是像 os.Args(0) 放置程序的名字。
	flag.Parse() //扫描参数列表（或者常量列表）并设置 flag 不需要排除第一个直接可用
	var s string = ""
	for i := 0; i < flag.NArg(); i++ { //返回参数的数量。解析后 flag 或常量就可用了。
		if i > 0 { //
			s += " "
			if *NewLine { // -n 被解析之后  flag值是 true
				s += Newline
			}
		}
		s += flag.Arg(i) //表示第 i 个参数 拼接参数
	}
	// 输出到标准输出1
	os.Stdout.WriteString(s)
}
