package main

import (
	"fmt"
	"runtime"
	"strconv"
)

/*一种可阅读性和打印性的输出。如果类型定义了 String() 方法，
它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。

————————————————
*/
type TwoInts struct {
	a int
	b int
}

func main() {
	// two1 := new(TwoInts)
	// two1.a = 12
	// two1.b = 10
	//short
	two1 := &TwoInts{12, 10}
	fmt.Printf("two1 is: %v\n", two1) //two1 is: (12/10)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1) //*main.TwoInts %T 会给出类型的完全规格
	// %#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）。
	fmt.Printf("two1 is: %#v\n", two1) //&main.TwoInts{a:12, b:10}
	// 上面的程序会给出已分配内存的总量
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024) //128kb
}
func (tn *TwoInts) String() string {
	// fmt.Sprintf("%v", t)  Sprintf方法会调用 string（）方法会造成死循环
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
