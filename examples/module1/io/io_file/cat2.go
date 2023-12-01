// 第二版
package main

import (
	"flag"
	"fmt"
	"os"
)

func cat2(f *os.File) {
	// 用切片读写文件
	const NBUF = 512
	var buf [NBUF]byte //数组
	for {
		// true是 条件一直成立 初始化nr, err
		// 把读取的内容存储到buf中
		// nr是读取的字节数
		switch nr, err := f.Read(buf[:]); true { //buf[:]
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			// os.Stdout.Write 返回写入的字节数 nr 和 nw应该相等
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func mainCat() {
	flag.Parse() // Scans the arg list and sets up flags
	if flag.NArg() == 0 {
		cat2(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat2(f)
		f.Close()
	}
}
