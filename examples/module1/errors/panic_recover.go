// panic_recover.go
package main

import (
	"fmt"
)

/*
Go panicking：

在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，
所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。
这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：
这个终止过程就是 panicking。

标准库中有许多包含 Must 前缀的函数，像 regexp.MustComplie 和 template.Must；
当正则表达式或模板中转入的转换字符串导致错误时，这些函数会 panic。

不能随意地用 panic 中止程序，必须尽力补救错误让程序能继续执行。

*/
func badCall() {
	panic("bad end") //抛出错误但是被恢复了
}

// defer-panic-recover
func test() {
	defer func() {
		// recover 只能在 defer 修饰的函数中使用
		// 用于取得 panic 调用中传递过来的错误值，如果是正常执行，调用 recover 会返回 nil 也就是没有执行 panic，且没有其它效果。
		// panic 会导致栈被展开直到 defer 修饰的 recover () 被调用或者程序中止。
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func mainR() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
