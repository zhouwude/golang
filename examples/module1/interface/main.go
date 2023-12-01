package main

import (
	"fmt"
)

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p *Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

type Array []int

func (a Array) getFirst() int {
	return a[0]
}
func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())
	s := []int{1, 2, 3, 4} //[]int 类型切片
	//getFirst定义在 Array上 即使类型不一样也不行
	// s.getFirst() s.getFirst undefined (type []int has no field or method
	// Array(s) 类型转换
	first := Array(s).getFirst()
	fmt.Println(first) //1

}
