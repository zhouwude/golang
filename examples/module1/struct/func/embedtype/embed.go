package main

import (
	"fmt"
	"math"
)

// 接口
/*
当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，
这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。
*/
type Engine interface {
	Start()
	Stop()
}

// 模拟经典面向对象语言中的子类和继承
//  Car自动获取到了 Engine的方法
type Car struct {
	Engine //匿名字段。。。
}

// 继承了 Engine的方法
// 当定义一个方法时，使用指针类型作为方法的接受者；
func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

type Point struct {
	x, y float64
}

/*
当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，
这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。
*/
func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// 可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。
// 直接作用域外层类型
// 当然类型可以有只作用于本身实例而不作用于内嵌 “父” 类型上的方法，
// func (n *NamedPoint) Abs() float64 {
// 	return n.Point.Abs() * 100.
// }

// 结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法。
// NamedPoint可以访问 Point 所有变量和方法在同一个包内
// 因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承
type NamedPoint struct {
	Point
	name string
}

// /类型中嵌入功能

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log //非内嵌匿名类型
}

// 内嵌类型
func main() {
	// c := new(Car)
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n)       //&{{3 4} Pythagoras} 指针
	fmt.Println(n.Abs()) // 打印5

	// ----------------
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())
	// 方法二
	c1 := &Customer1{Name: "Barak Obama", Log1: Log1{"1 - Yes we can!"}}
	c1.Add("2 - After me the world will be a better place!")
	fmt.Println(c1)

	// Log:{1 - Yes we can!
	// 2 - After me the world will be a better place!}

}

// 聚合（或组合）：包含一个所需功能类型的具名字段。
func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

//  String()方法 应该是字面量表达的方法吧 ，类似于 NSObject的 description方法打印结构的默认方法
func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

// 内嵌：内嵌（匿名地）所需功能类型
type Log1 struct {
	msg string
}

type Customer1 struct {
	Name string
	Log1 //内嵌匿名类型
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

// 自定义了输出方式
func (l *Log1) String() string {
	return l.msg
}

func (c *Customer1) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log1)
}
