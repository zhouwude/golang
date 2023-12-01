package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

//怎么支持 go test 扫描了所有_test.go 为结尾的文件 将正式代码和测试代码放在同一个目录
// foo.go的测试代码放在 foo_test.go中
// go test path 在指定目录下面test
func TestIncrease(t *testing.T) {
	t.Log("Start testing")
	result := add(1, 2)
	// 断言
	fmt.Println(assert.Equal(t, result, 3))
}
