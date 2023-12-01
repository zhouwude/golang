package main

import (
	"fmt"
	"sort"
)

//空接口或者最小接口 不包含任何方法，它对实现不做任何要求：
// 相当于其他语言的 Any 类型任何类型都实现了空接口iOS中的 NSobject 协议就是
// ******任何其他类型都实现了空接口
var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

// 每个 interface {} 变量在内存中占据两个字长：
// 一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针。
type Any interface{}
type specialString string

//
var whatIsThis specialString = "hello"

func TypeSwitch() {
	// 匿名函数
	// interface{}就是一个空接口
	//任何人一个人类型都可以当做 testFunc的参数
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		//
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}

// 之前的切片只能一种类型
// 构建通用类型或包含不同类型变量的数组
type Element interface{}
type Vector struct {
	a []Element //切片类型 实现了空接口的类型都可以
}

//接收者是指针 类型是指针才可以
func (p *Vector) At(i int) Element {
	return p.a[i]
}
func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

type SqrInterface interface {
	Sqr() float32
}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
	/*
			val has the value: 5
		val has the value: ABC
		val has the value: &{Rob Pike 55}
		Type pointer to Person *main.Person*/
	TypeSwitch()
	//any hello is a special String!
	fmt.Println(sort.SearchInts([]int{1, 2, 3, 4, 5}, 5)) //下标是4
	// 复制数据切片至空接口切片不能直接赋值 需要使用 for range一个个赋值

}
