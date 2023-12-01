// parse.go
package parset

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseError表示在将单词转换为整数时出现错误。
type ParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Err   error  // 导致这个错误的原始错误(如果有的话)。
}

// String returns a human-readable error message.
// 自定义输出的内容 格式
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			// r 是 interface{} 接口断言
			err, ok = r.(error)
			if !ok {
				//返回一个 error
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input) //通过空格分割
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field) //str -》 num
		if err != nil {
			// 指针
			// 字面量写法取地址是指针 &var
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
