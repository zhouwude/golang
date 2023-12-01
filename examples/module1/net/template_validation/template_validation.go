package main

import (
	"fmt"
	"text/template"
)

// 模板验证
func main() {
	tOk := template.New("ok")
	//检查模板的语法是否定义正确，对 Parse 的结果执行 Must 函数
	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	tErr := template.New("error_template")
	template.Must(tErr.Parse(" some static text {{ .Name }")) //少一个}
}
