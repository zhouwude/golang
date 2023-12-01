// simple_webserver.go
package main

import (
	"io"
	"net/http"
)

// 当您将action属性设置为#时，表单数据将被提交到当前页面的URL。
const form = `<html><body><form action="#" method="post" name="bar">
		      <input type="text" name="in"/>
			  <input type="submit" value="Submit"/>
			  </form></html></body>`

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	//写入数据 response 写入
	//把内容写入 response有很多种方法 这里只是一种
	io.WriteString(w, "<h1>hello, world</h1>")
}

/* handle a form, both the GET which displays the form
   and the POST which processes it.*/
func FormServer(w http.ResponseWriter, request *http.Request) {
	//设置头
	/*
				更安全的做法是在处理器中使用 w.Header().Set("Content-Type", "text/html") 在写入返回之前将 header 的 content-type 设置为 text/html

			content-type 会让浏览器认为它可以使用函数 http.DetectContentType([]byte(form)) 来处理收到的数据

	;*/
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		//处理表单数据，注意ParseForm必须在提取表单数据  之前调用
		// request.ParseForm ();
		// io。WriteString (w, request.Form[的][0])
		// 使用 request.FormValue("inp") 通过文本框的 name 属性 inp 来获取内容，并写回浏览器页面。
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
