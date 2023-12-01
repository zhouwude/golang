package main

import (
	"fmt"
	"io"
	"reflect"

	duck "github.com/cncamp/golang/examples/module1/reflect/interfaceanddynamic/duck_dance"
)

/* ******任何提供了接口方法实现代码的类型都隐式地实现了该接口，而不用显式地声明。
和其它语言相比，Go 是唯一结合了接口值，静态类型检查（是否该类型实现了某个接口），
运行时动态转换的语言，
并且不需要显式地声明类型是否满足某个接口。
该特性允许我们在不改变已有的代码的情况下定义和使用新接口。
接收一个（或多个）接口类型作为参数的函数，其实参可以是任何实现了该接口的类型。
 实现了某个接口的类型可以被传给任何以此接口为参数的函数 。
*/
type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Rank() int {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Rank() int {
	return 2
}

// 显式地指明类型实现了某个接口#
type Fooer interface {
	Foo()
	ImplementsFooer()
}
type Bar struct{}

func (b Bar) ImplementsFooer() {}
func (b Bar) Foo()             {}
func main() {
	b := new(duck.Bird)
	//不需要转换类型 b实现了接口方法
	duck.DuckDance(b)
	/*I am walking!
	I am quacking!
	I am walking!
	I am quacking!
	I am walking!*/

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}
	/*Looping through shapes for area ...
	Shape details:  {5 3}
	Area of this shape is:  15
	Shape details:  &{5}
	Area of this shape is:  25
	Looping through topgen for rank ...
	Shape details:  {5 3}
	Topological Genus of this shape is:  2
	Shape details:  &{5}
	Topological Genus of this shape is:  1*/
	// 确定大小的就是数组 值类型 没有就是切片 引用类型
	fmt.Println(reflect.TypeOf([5]int{1, 2, 3, 4, 5}).Kind()) //array
	fmt.Println(reflect.TypeOf([]int{1, 2, 3, 4, 5}).Kind())  //slice
}

type xmlWriter interface {
	WriteXML(w io.Writer) error
}

// Exported XML streaming function.
func StreamXML(v interface{}, w io.Writer) error {
	// v是一个空接口不确定实现了其他接口 线动态判断
	//v是否实现了xmlWriter的方法
	if xw, ok := v.(xmlWriter); ok {
		// It’s an  xmlWriter, use method of asserted type.
		return xw.WriteXML(w)
	}
	// No implementation, so we have to use our own function (with perhaps reflection):
	return encodeToXML(v, w)
}

// Internal XML encoding function.

func encodeToXML(v interface{}, w io.Writer) error {
	// ...
	return io.ErrClosedPipe
}
