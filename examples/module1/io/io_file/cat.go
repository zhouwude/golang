// 用 buffer 读取文件
package main

// buffer缓冲区
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //读到换行符停止
		if err == io.EOF {            //io.EOF 文件结尾
			break
		}
		//输出到命令行标准输出C
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func mainC() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
