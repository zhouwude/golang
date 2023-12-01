package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

// 这个程序是一个 web 服务器，所以它必须在命令行启动（译者注：不要在 IDE 中启动，否则会找不到路径，必须在命令行启动）
const lenPath = len("/view/")

// MustCompile类似于Compile，但是如果表达式不能被解析就会出现panic。
// 它简化了保存编译正则表达式的全局变量的安全初始化。 保证不会出错用这个方法
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

//map
var templates = make(map[string]*template.Template)
var err error

//
type Page struct {
	Title string
	Body  []byte //字节切片
}

//包的初始化
func init() {
	for _, tmpl := range []string{"edit", "view"} {
		/*
					   template.Must(template.ParseFiles(tmpl + ".html"))
					   函数将模板文件转换成一个 *template.Template （Template 结构体的指针），
					   为了提高效率，我们只在我们的程序中转换一次，放在 init() 函数中就可以很方便的实现了。
					   这个模板对象被保存在内存中的一个以 html 文件名称为索引的 map 中。

			         ******这就是模板缓存 推荐用法
		*/
		//函数参数两个参数 正好ParseFiles 返回两个参数 可以这样直接调用 nice
		templates[tmpl] = template.Must(template.ParseFiles(tmpl + ".html"))
	}
}

//*******这里必须先 go buid wiki.go 在执行可执行文件 要不然会报错
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	err := http.ListenAndServe("localhost:8089", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		title := r.URL.Path[lenPath:]
		// fmt.Fprintf(w, title)
		// return
		// a-zA-Z0-9 验证title 有效
		if !titleValidator.MatchString(title) {
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil { // page not found
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := load(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body") //表单中的 body值
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// 模板和结构体输出到页面
	// 并且写入到 ResponseWriter w 中 返回给客户端
	//将内容写到 response
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// file created with read-write permissions for the current user only
	//全部写入 0600 0b110000000给了读写权限
	//写带文件
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	//全部读完内容
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
