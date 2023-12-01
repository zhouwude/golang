package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	// 还可以在你的代码中使用一些预定义的模板函数，
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
	// hello Mary!
}
