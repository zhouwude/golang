//Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Hello world, web服务器
// hello-requests
var helloRequests = expvar.NewInt("hello-requests")

// 开始对常量、变量和类型的定义或声明
// flags: 默认值
var webroot = flag.String("root", "/Users/zhouwude", "web root directory")

// 简单标志服务器 bool 类型 只要解析到 就是true
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

// Simple counter server. POSTing to it will set the value.
type Counter struct {
	n int
}

// 通道类型
type Chan chan int

func main() {
	flag.Parse() //解析参数
	// Handler 接口类型方法是 ServeHTTP只要是实现了 ServeHTTP方法的类型默认实现了该接口
	// type HandlerFunc func(ResponseWriter, *Request) 是函数类型的别名 alias
	// 它是一个可以把普通的函数当做 HTTP 处理器的适配器。
	// 如果 f 函数声明的合适，HandlerFunc(f) 就是一个执行了 f 函数的处理器对象。
	// HandlerFunc(func)相当于类型转换 把一个一样的函数转成HandlerFunc格式 因为HandlerFunc实现了 Handle接口
	// HandlerFunc是一个类型 而不是一个方法
	http.Handle("/", http.HandlerFunc(Logger))
	http.Handle("/go/hello", http.HandlerFunc(HelloServer))
	// The counter is published as a variable directly.
	ctr := new(Counter) //new 一个结构体
	// 它可以创建一个变量（可能是 int、float 或者 string 类型），并且通过发布去公开他们
	// 使用 JSON 格式在 HTTP /debug/vars （译者注：就是可以通过 localhost:12345/debug/vars
	// 浏览器中 localhost:12345/debug/vars查看这些被公开的变量）
	// 公开这些变量。它一般用于服务器中的操作计数器
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	// http.Handle("/go/", http.FileServer(http.Dir("/tmp"))) // uses the OS filesystem
	// FileServer 返回一个 root 参数的值为根目录的文件来处理 HTTP 请求。通过 http.Dir 去使用操作系统的文件系统
	// 可以在 /tmp 目录下创建一个 ggg.html , 再访问 /go/ggg.html 的时候就会直接在浏览器中显示 ggg.html 的内容。
	// *取内容
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags", http.HandlerFunc(FlagServer))
	http.Handle("/args", http.HandlerFunc(ArgServer))
	http.Handle("/chan", ChanCreate())
	http.Handle("/date", http.HandlerFunc(DateServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}

func Logger(w http.ResponseWriter, req *http.Request) {
	// Logger 会用 w.WriteHeader (404) 记录一个 404 Not Found header。

	// 这个技术通常很有用，当 web 处理代码发生错误的时候，

	log.Print(req.URL.String()) //打印日志
	w.WriteHeader(404)
	// 字符串 -》 字符切片
	w.Write([]byte("oops"))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	// hello-requests
	// 没调用一次 hello-requests +1
	helloRequests.Add(1)
	io.WriteString(w, "hello, world!\n")
}

//这使得Counter满足expvar。Var接口，以便我们可以导出
//直接调用。
func (ctr *Counter) String() string { return fmt.Sprintf("%d", ctr.n) }

// 实现 http调用接口
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET": // increment n
		ctr.n++ //只能放在后面
	case "POST": // set n to posted value
		buf := new(bytes.Buffer) //pointer
		io.Copy(buf, req.Body)   //复制
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	//response
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func FlagServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Flags:\n")
	// 遍历所有的标志 没有指定值 就是默认值
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

// simple argument server
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

// 通道的工厂模板：以下函数是一个通道工厂，
// 启动一个匿名函数作为协程以生产通道：
func ChanCreate() Chan {
	// 不带缓冲的 chan
	c := make(Chan)
	// 匿名的立即执行函数

	go func(c Chan) {
		// 服务端通过死循环来从 chan *Request 接收请求
		// ，为了避免长时间运行而导致阻塞，可以为每个请求都开一个 goroutine 来处理：
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

// 实现了接口类型的方法 就是隐式的实现了 Handle 接口
// 没调用一次值会变化
func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 如果没有可用的发送者 则一直阻塞
	io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}
func (ch Chan) ChanResponse(w http.ResponseWriter, req *http.Request) {
	timeout := make(chan bool)

	go func() {
		time.Sleep(30e9) //30s
		timeout <- true  //有效的发送者
	}()

	select {
	case msg := <-ch:
		io.WriteString(w, fmt.Sprintf("channel send #%d\n", msg))
	case <-timeout:
		return
	}
}

// exec a program, redirecting output
func DateServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Pipe返回一对连接的文件;从r读取写入w的返回字节。
	//返回文件和错误(如果有的话)。
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}
	// 启动执行命令
	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
	defer r.Close()
	w.Close()
	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}
