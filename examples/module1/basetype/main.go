package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
	"unicode"
)

func main() {
	var _ int8 = 10
	// int8 -128 - 127
	/*  (现在的计算机用补码存储整数数值 第一位永远是符号为 0 为正数 1为负数)****
		原码，就是用二进制表示的原始编码，
		反码，就是除符号位外，其它位取反
		补码，正数的补码等于原码，负数的补码等于反码加1

		在这里以byte为例，
	计算机中1byte=8bit，因此byte的最大正数为二进制的01111111 = 2^7-1 = 127 （0位符号位），这个没问题，
	对于负数而言，其补码由于需要反码加1 ， 对于11111111其补码就是10000001=-1，
	依次类推11111110补码为10000010=-2........一直到10000001补码为11111111=-127，
	接下来最大的负数为10000000十进制其实是-0 ，其实是没有意义的（与+0重复），
	计算机就将这一数值10000000表示为最大负数-128
	10000000 -> 补码 （11111111 + 1) -> 10000000 = -0 (超过八位 舍去一位)
	会出现一个+0和一个-0。印度人他们规定-0为-128，这样就与计算机的补码（程序都是按补码运行的）完美的结合在一起。
	*/
	// 11111111 这是负数 反码-》10000000 补码 10000001 就是负一
	// 进制转换
	fmt.Println(strconv.FormatBool(false))
	//%v默认打印方式
	// 你可以通过增加前缀 0 来表示 8 进制数（如：077），增加前缀 0x 来表示 16 进制数（如：0xFF），
	fmt.Println(strconv.FormatInt(10, 2))          //1010
	fmt.Println(strconv.FormatInt(0b11111111, 10)) //255
	fmt.Println(strconv.FormatInt(0b01111111, 10))
	fmt.Println(strconv.FormatInt(0x768, 10)) //1896 16进制
	fmt.Println(strconv.FormatInt(0o365, 10)) //245 八进制
	fmt.Println(strconv.FormatInt(0365, 10))  //245
	// 把其他进制转换为10进制
	fmt.Println(strconv.ParseInt("1010101", 2, 64)) //85
	fmt.Println(math.MaxInt8)                       //127
	a, b := math.Modf(2.34)
	fmt.Println(a, b) //2 0.33999999999999986
	a++
	fmt.Println(a) //自增和自减运算符都只能放在后面 只能作为表达式不能作为语句
	// c := a++ 错误写法
	// func (a++) 错误写法
	// 随机数
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	fmt.Println(timens) //517535000
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", rand.Float32())
	}
	// byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，
	// 完全没有问题。例如：var ch byte = 'A'；字符使用单引号括起来。

	c := 'a'
	fmt.Println()
	fmt.Println("c: ", c) //c:  97
	fmt.Printf("%c", c)   //a
	fmt.Println()
	fmt.Println(string(c)) //a
	// var ch byte = 65
	// var ch byte = '\x41'
	// 2的四次方等于16所以16进制等于个四个二进制 所以一个字节可以这样表示\x14
	// 2的三次方等于8所以8进制等于3个二进制 \445一个字节
	fmt.Println(int('周')) //21608
	// 因为 Unicode 至少占用 2 个字节，所以我们使用 int16
	// 或者 int 类型来表示。如果需要使用到 4 字节，则会加上 \U 前缀；
	// 前缀 \u 则总是紧跟着长度为 4 的 16 进制数，前缀 \U 紧跟着长度为 8 的 16 进制数。
	// \u00ff 两个字节 \U0000004f 四个字节
	fmt.Println(string(rune(21006)))           //刎
	var ch int = '\u0041'                      //2byte
	var ch2 int = '\u03B2'                     //2byte
	var ch3 int = '\U00101234'                 //4byte
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	fmt.Println()
	fmt.Println("isdigit ", unicode.IsDigit(32))
	// byte类型就是一个数字
	fmt.Println(unicode.IsUpper(65))
	//switch 中加语句的时候一定要加逗号;
	switch s1 := 1 + 3; {
	case s1 > 1:
		fmt.Println(s1)
	case s1 < 0:
		fmt.Println(s1)
	}
}
