package main

import (
	"fmt"
	"math"
	"time"

	person1 "github.com/cncamp/golang/examples/module1/struct/func/person"
)

type TwoInts struct {
	a int
	b int
}
type TwoInts1 struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	// two2.AddThem 可以替代 (&two2).AddThem() 编译器自动处理了（自动解引用）
	/*可以有连接到类型的方法，也可以有连接到类型指针的方法。

	但是这没关系：对于类型 T，如果在 *T 上存在方法 Meth()，
	并且 t 是这个类型的变量，那么 t.Meth() 会被自动转换为 (&t).Meth()。
	*/
	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	s := make(IntVector, 0)
	s = append(s, IntVector{1, 2, 3}...)

	fmt.Println(s.Sum()) //6

	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())

	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())    //{1}
	fmt.Println("write: ", b1) //write:  {1}  write接收者是b1的拷贝所以B1并没有改变
	fmt.Println(b1.writeP())
	fmt.Println("writeP: ", b1) //writeP:  {2}
	b2 := new(B)                // b2是指针
	b2.change()
	fmt.Println(b2.write()) //{1}

	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
	fmt.Println()
	p1 := new(person1.Person)
	p1.SetFirstName("Eric")
	fmt.Println(p1.FirstName()) // Output: Eric
}

/*
receiver_type
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
如果 recv 是 receiver_type methodName 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.methodName()。

*******如果recv 一个指针 *receiver_type会自动解引用。************

如果方法不需要使用 recv 的值，可以用 _ 替换它

***类型的代码和绑定在它上面的方法的代码可以不放置在一起，
它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。
*/
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}
func (tn *TwoInts1) AddThem() int {
	return tn.a + tn.b
}

// 这里接收者不一样方法可以重载否则 无法重载。
//********* 如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。
// 否则，就在普通的值类型上定义方法。
// 如果接受者是值属于拷贝没法改变接收者的数据用指针最好。
// 因为值类型一直是传递拷贝
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
func (tn *TwoInts1) AddToParam(param int) int {
	return tn.a + tn.b + param
}

type IntVector []int

// 这是方法不是函数方法 方法跟绑定类型可以不在一起但一定要在一个包里面。
// 非结构体方法 接收者是一个切片类型
// 命名返回值可以直接参与计算 很好的习惯 还可以唉defer 中修改返回值。
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

//可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。
// 或者像下面这样将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
type myTime struct {
	time.Time //anonymous xxxx
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

type B struct {
	thing int
}

// 指针作为接收者
func (b *B) change() { b.thing = 1 }

//值作为接收者 这里 b 是一个拷贝,,,原来的值并不会被修改。。。
// 仅当一个结构体的方法想改变结构体本身时，使用结构体指针作为方法的接受者，否则使用一个结构体值类型
func (b B) write() string {
	b.thing = 2 //改变值
	return fmt.Sprint(b)
}
func (b *B) writeP() string {
	b.thing = 2 //改变值
	return fmt.Sprint(b)
}

type Point3 struct{ x, y, z float64 }

// A method on Point3
func (p Point3) Abs() float64 {
	// 只需要 计算结果

	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }
