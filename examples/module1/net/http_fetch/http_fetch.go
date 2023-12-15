// httpfetch.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// 在下边这个程序中，数组中的 url 都将被访问：会发送一个简单的 http.Head() 请求查看返回值；它的声明如下：func Head(url string) (r *Response, err error)

func main() {

	// get方法
	res, err := http.Get("http://www.google.com")
	CheckError(err)
	// ioutil将整个文件的内容读到或者写到一个文件里
	data, err := io.ReadAll(res.Body)
	CheckError(err)
	fmt.Printf("Got: %q", string(data))
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
}
