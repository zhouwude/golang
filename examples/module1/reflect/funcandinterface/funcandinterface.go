package main

import (
	"fmt"
)

/*
*********在接口上调用方法时****在类型上调用就没事***，必须有和方法定义时相同的接收者类型
或者是可以从具体类型 P 直接可以辨识的：

指针方法可以通过指针调用
值方法可以通过值调用
接收者是值的方法可以通过指针调用，因为指针会首先被解引用
*********接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址*********
将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

译注

Go 语言规范定义了接口方法集的调用规则：

类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
类型 T 的可调用方法集包含接受者为 T 的所有方法
类型 T 的可调用方法集不包含接受者为 *T 的方法
*/
type List []int

func (l List) Len() int {
	return len(l)
}

// 要改变接收者的数据必须是指针 否则是一个拷贝
func (l *List) Append(val int) {
	*l = append(*l, val)
}

// interface{}
type Appender interface {
	Append(int)
}

// 隐私转换成了接口类型Appender
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	//在 lst 上调用 CountInto 时会导致一个编译器错误，
	// 因为 CountInto 需要一个 Appender，而它的方法 Append 只定义在指针上
	// *****CountInto(lst, 1, 10)
	CountInto(&lst, 1, 10) //接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
	if LongEnough(lst) {   // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}

}
