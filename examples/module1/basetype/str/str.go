package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	/*
		当需要对一个字符串进行频繁的操作时，谨记在 go 语言中字符串是不可变的（类似 java 和 c#）。
		使用诸如 a += b 形式连接字符串效率低下，尤其在一个循环内部使用这种形式。
		这会导致大量的内存开销和拷贝。应该使用一个字符数组代替字符串，将字符串内容写入一个缓存中。
		例如以下的代码示例：

	*/
	// var b bytes
	// for condition {
	// 	b.WriteString(str) // 将字符串str写入缓存buffer
	// }
	var b bytes.Buffer
	for i := 0; i < 10; i++ {
		b.WriteString("hello world")
	}
	// Buffer 转成 字符串-----
	fmt.Println(b.String())

	fmt.Println(int('😵'))                      //128565
	fmt.Println(strconv.FormatInt(128565, 16)) //1f635
	// 都是字节来的
	fmt.Println("周武德") //周武德
	// 这种转换方案只对纯 ASCII 码的字符串有效。
	fmt.Println("周武德"[1]) //145 第二个字节是145
	fmt.Println(strings.Join([]string{"1", "1"}, ""))
	// 如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位：
	str1 := "我是周五的"
	fmt.Println(strings.IndexRune(str1, '的')) //12
	// 字符串的长度
	fmt.Println(utf8.RuneCountInString("我是周五的")) //5
	fmt.Println(len([]rune(str1)))               //5
	fmt.Println(len(str1))                       //15
	fmt.Println(strings.Repeat("d==", 4))
	fmt.Println(strings.Fields("东南王气多 当我看到清楚 大家我带你玩"))
	// [东南王气多 当我看到清楚 大家我带你玩] 利用空白符分割
	fmt.Println(strings.Split("dwdqedqq", "d")) //[ w qe qq]

	// 下标都是根据字节走的 中文字符三个字节下标是 0 3 6 9 12
	for index, value := range str1 {
		fmt.Println(index, string(value))
	}
	reader := strings.NewReader(str1)
	bytes, _, _ := reader.ReadRune()
	fmt.Println("----", bytes) //25105
	// 该包包含了一些变量用于获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize。
	fmt.Println(strconv.IntSize) //64

	// 数字转字符串
	fmt.Println(strconv.Itoa(20016))
	fmt.Println(strconv.Atoi("10"))     //10
	fmt.Println(unicode.ToLower(20016)) //
	// 码位转字符串
	fmt.Println(string(rune(128565))) //😵

	ss := "zhouwude"
	ss1 := ss[3:6]
	fmt.Printf("--%p--%p", &ss, &ss1) //--0xc000096310--0xc000096320

	fmt.Println([]byte("周武德")) //[229 145 168 230 173 166 229 190 183]
	// 如何获取一个字符串的字节数：len(str)
	// 如何获取一个字符串的字符数：
	fmt.Println(len([]rune("周武德"))) //3

}
