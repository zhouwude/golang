package main

import (
	"fmt"
	"os"
	"text/template"
)

// 字段替代
type Person struct {
	Name                string
	nonExportedAgeField string //小写开头不能导出 会报错 必须大写
}

func main() {
	// New用给定的名称分配一个新的未定义模板。
	t := template.New("hello")
	t1 := template.New("hello1")
	//如果 Name 是一个结构体中的字段，并且它的值需要在合并时替换，
	// 那么在模板中包含文本 {{.Name}} ；当 Name 是一个 map 的索引时，也可以这样使用。
	t, _ = t.Parse("hello {{.Name}}!")
	// 但是直接使用 {{ . }} ，不管字段是否可以导出，会将两个字段全部输出。即使是小写
	t1, _ = t1.Parse("hello {{.}}!")
	// 当参数是一个定义好的模板文件的路径时
	// t, _ = t.ParseFiles("")
	p := Person{Name: "Mary", nonExportedAgeField: "31"}
	// 把结果输出到Stdout 、、hello Mary!
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	if err := t1.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
	// hello {Mary 31}!
}

// Output:   hello Mary!
