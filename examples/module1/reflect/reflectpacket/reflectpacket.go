package main

import (
	"fmt"
	"reflect"
)

/*
用于修改某些操作的默认行为，
等同于在语言层面做出修改，
所以属于一种“元编程”（meta programming），
即对编程语言进行编程。
*/
//枚举
const (
	Invalid reflect.Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

type T struct {
	A int
	B string
}

func main() {

	var x float64 = 3.4
	fmt.Println(reflect.TypeOf(x))  //float64
	fmt.Println(reflect.ValueOf(x)) //3.4
	type MyInt int
	var m MyInt = 5

	v := reflect.ValueOf(m) //类型
	fmt.Println(v.Kind())   //int
	// 变量 v 的 Interface() 方法可以得到还原（接口）值，
	// 所以可以这样打印 v 的值：fmt.Println(v.Interface())
	// 任何变量都遵循空接口
	fmt.Println(v.Interface()) //5

	// ***********函数通过传递一个 x 拷贝创建了 v ***
	// 也就是说改变 V 的值原始的 x 的值并不会变化
	v = reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                //type: float64
	fmt.Println("kind:", v.Kind())                //kind: float64
	fmt.Println("value:", v.Float())              //value: 3.4 返回这个 float64 类型的实际值
	fmt.Println(v.Interface())                    //3.4
	fmt.Printf("value is %5.2e\n", v.Interface()) //value is 3.40e+00
	//判断接口变量的类型
	y, ok := v.Interface().(float64)
	fmt.Println("ok is", ok)                     //ok is true
	fmt.Println(y)                               //3.4
	fmt.Println("settability of v:", v.CanSet()) //settability of v: false

	v = reflect.ValueOf(&x)             //传递地址 引用类型 可以修改原始值
	fmt.Println("type of v:", v.Type()) //type of v: *float64
	// 通过 Type () 我们看到 v 现在的类型是 *float64 并且仍然是不可设置的。
	fmt.Println("settability of v:", v.CanSet()) //settability of v: false
	// 要想让其可设置 不管是指针还是原始类型都需要做这个操作
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)         //The Elem of v is:  3.4
	fmt.Println("settability of v:", v.CanSet()) //settability of v: true
	v.SetFloat(3.1415)                           // this works!
	fmt.Println(v.Interface())                   //3.1415
	fmt.Println(v)                               // 3.1415
	// 解构体

	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ) //main.NotknownType
	knd := value.Kind()
	fmt.Println(knd) // struct

	// iterate through the fields of the struct:
	// NumField() 方法返回结构体内的字段数量
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}
	/*
			Field 0: Ada
		Field 1: Go
		Field 2: Oberon*/
	// call the first method, which is String():
	// 一个切片在未初始化之前默认为 nil 引用类型未初始化就是 nil 而值类型是该类型对应的零值。
	results := value.Method(0).Call(nil)
	fmt.Println(value.Method(0).Type()) //func() string
	fmt.Println(value.Method(0).Kind()) //func
	fmt.Println(results)                // [Ada - Go - Oberon]
	// e 2

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	/*
			0: A int = 23
		1: B string = skidoo
	*/
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t) //t is now {77 Sunset Strip}

}
