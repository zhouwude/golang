// template_with_end.go
package main

import (
	"os"
	"text/template"
)

// 字段替代
type Person struct {
	Name                string
	nonExportedAgeField string //小写开头不能导出 会报错 必须大写
}

func main() {
	// with 语句将点的值设置为管道的值。如果管道是空的，就会跳过 with 到 end 之前的任何内容；
	// 当嵌套使用时，点会从最近的范围取值：
	p := &Person{Name: "zhouwude", nonExportedAgeField: "ss"}
	t := template.New("test")
	t, _ = t.Parse("{{with .Name}}{{.}}{{end}}!\n")
	t.Execute(os.Stdout, p)
	// zhouwude!

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t.Execute(os.Stdout, nil)
}

/* Output:
hello!
hello Mary!
*/
