package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	// r.URL.Path[1:] 组成的字符串后边的 [1:] 表示 “创建一个从第一个字符到结尾的子切片”，用来丢弃掉路径开头的 “/”
	// 完成了本次写入
	//讲内容写入到 ResponseWriter也就是 请求者Request能收到的内容
	// Path这里不包含域名和端口 只有 Path部分
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func main() {
	// http.Handle("/", http.HandlerFunc(HelloServer))
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	// http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))前两行（没有错误处理代码）可以替换成以下写法：
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
