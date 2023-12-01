package main

import (
	"fmt"

	"github.com/cncamp/golang/examples/module1/errors/even"
)

func mainM() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
