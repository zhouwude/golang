package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	// 只有类型  没有变量名 类型当成变量名 在一个结构体中对于每一种数据类型只能有一个匿名字段
	int    // anonymous field
	innerS //anonymous field
}
type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	//直接进入内层
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
	b := B{A: A{1, 2}, bx: 3.0, by: 4.0}
	//直接取内层的属性
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	// 如果有相同变量名 则不能省略
	fmt.Println(b.A.ax, b.A.ay, b.bx, b.by)
	fmt.Println(b.A)
}
