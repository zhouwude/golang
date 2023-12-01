// panic_package.go
package main

import (
	"fmt"

	parset "github.com/cncamp/golang/examples/module1/errors/parse"
)

/*
这是所有自定义包实现者应该遵守的最佳实践：

1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic ()

2）向包的调用者返回错误值（而不是 panic）。
*/
func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := parset.Parse(ex)
		if err != nil {
			fmt.Println(err) // here String() method from ParseError is used
			continue
		}
		fmt.Println(nums)
	}
	/*
			Parsing "1 2 3 4 5":
		  [1 2 3 4 5]
		Parsing "100 50 25 12.5 6.25":
		  pkg: pkg parse: error parsing "12.5" as int
		Parsing "2 + 2 = 4":
		  pkg: pkg parse: error parsing "+" as int
		Parsing "1st class":
		  pkg: pkg parse: error parsing "1st" as int
		Parsing "":
		  pkg: no words to parse*/
}
