package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Bar struct {
	thingOne string
	thingTwo int
}
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

//读取标签
func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	fmt.Println(unsafe.Sizeof(File{}))          //24
	fmt.Println(unsafe.Sizeof(&File{}))         //8 指针类型 本身是一个变量， 只占用八个字节
	fmt.Println(unsafe.Sizeof(new(File)))       //8
	fmt.Println(unsafe.Sizeof(File{8, "----"})) //24
	type Foo map[string]string
	fmt.Println(unsafe.Sizeof(make(Foo))) // 引用类型 都是8

	y := new(Bar) //指针
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	// 直接
	y.thingOne = "hello"
	y.thingTwo = 1

	// NOT OK
	// z := make(Bar) // 编译错误：cannot make type Bar
	// (*z).thingOne = "hello"
	// (*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	// new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
	// u := new(Foo)
	// (*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	// (*u)["y"] = "world"
	//字符串中包含字符串
	fmt.Println(`my name is "周武德"`) //my name is "周武德"
	tt := TagType{true, "Barak Obama", 1}

	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}
