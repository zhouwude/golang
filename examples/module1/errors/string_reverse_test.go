// string_reverse_test.go
package main

import (
	"testing"

	strev "github.com/cncamp/golang/examples/module1/errors/string_reverse"
)

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest{
	{"ABCD", "DCBA"},
	{"CVO-AZ", "ZA-OVC"},
	{"Hello 世界", "界世 olleH"},
}

// 测试程序必须属于被测试的包，***并且文件名满足这种形式 *_test.go，所以测试代码和包中的业务代码是分开的。
// *******测试函数必须以Test大写开头后面是测试函数的字母描述， 例如Reverse反转
func TestReverse(t *testing.T) {
	/*
		in := "CVO-AZ"
		out := Reverse(in)
		exp := "ZA-OVC"
		if out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", in, exp, out)
		}
	*/
	// testing with a battery of testdata:
	for _, r := range ReverseTests {
		exp := strev.Reverse(r.in)
		if r.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", r.in, exp, r.out)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		strev.Reverse(s)
	}
}
