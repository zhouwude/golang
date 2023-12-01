package main

import (
	"fmt"
	"strconv"
)

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

// 内嵌两个匿名类型
type CameraPhone struct {
	Camera
	Phone
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

// 接收者和内嵌匿名类型有同样的方法
func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

// 组件编程
type Integer int

func (i *Integer) String() string {
	return strconv.Itoa(int(*i))
}

// 多重继承
func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
	v := new(Voodoo)
	v.Magic() //先调用自己的方法 没有再找匿名类型方法
	v.MoreMagic()
	// voodoo magic
	// base magic
	// base magic

}
