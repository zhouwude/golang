package main

import (
	"fmt"
	"text/template"
)

// 模板验证
func main() {
	tOk := template.New("ok")
	//检查模板的语法是否定义正确，对 Parse 的结果执行 Must 函数
	// template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	validTemplate(tOk)
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	tErr := template.New("error_template")
	validTemplate(tErr)
	// template.Must(tErr.Parse(" some static text {{ .Name }")) //少一个}
}
func validTemplate(tem *template.Template) {
	defer func() {

		/*		// recover 只能在 defer 修饰的函数中使用
				// 用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil 也就是没有执行 panic，且没有其它效果。
				// panic 会导致栈被展开直到 defer 修饰的 recover () 被调用或者程序中止。*/
		if e := recover(); e != nil {
			fmt.Println(e) //template: error_template:1: unexpected "}" in operand
		}
	}()
	template.Must(tem.Parse(" some static text {{ .Name }"))
}
