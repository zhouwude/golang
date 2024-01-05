package main //package main 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/cncamp/golang/examples/module1/struct/matrix"
	p "github.com/cncamp/golang/examples/module1/struct/person"
)

// alias 类型转换作用很大
type (
	IZ  int
	FZ  float64
	STR string
)

// 带标签的解构体
type MyType struct {
	Name string `json:"name"` //annotation 注解kubernetes常规操作
	Address
}
type Address struct {
	City string `json:"city"`
}

func main() {
	mt := MyType{Name: "test", Address: Address{City: "shanghai"}}
	b, _ := json.Marshal(&mt)
	fmt.Println(string(b))
	mt1 := MyType{}
	fmt.Println(mt)
	// {test {shanghai}}
	// //{{} 0 0} 解构体初始化不需要一定给变量赋值不赋值默认给类型的零值
	fmt.Println(mt1) //{ {}} 都是空字符串
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	println("tag------------", tag)
	// return
	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	//可以直接访问 TypeA.P1
	println(tb.P1)
	fmt.Println(test2(100))
	m := make(map[string]int)
	m["a"] = 1
	// init的方法 在别的语言 中 叫做构造器或者初始化方法
	m["b"] = 2
	if v, success := m["b"]; success {
		fmt.Println(v)
	}
	fmt.Println(m)
	var s1 S1 //s1在这里就会分配内寸，属性的值就是类型对应的零值。
	s1.a = "a"
	// 只要变量申明了就会赋予 变量类型的零值

	s2 := S1{a: "a"}
	fmt.Println(s2)
	// 使用 new 返回的是指针引用类型 不使用的话是值类型
	s3 := new(S1)
	fmt.Println("s3.a s3: ", s3.a)
	fmt.Println("s3: ", s3) //指针类型
	fmt.Println("s3: ", *s3)
	// %v 标准打印不知道具体类型的时候可以使用

	typeB := TypeB{"p2", TypeA{"p1"}}                //顺序一致 不可省略
	typeC := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}} // 不必一致可以省略某些属性 默认零值。
	fmt.Printf("format1 %v \n", typeB)
	fmt.Printf("format2 %v \n", typeC)
	fmt.Println("----------------------------")
	p.PersonF()

	matrix.NewMatrix(10)
	// m := new(matrix.matrix)
	// make 只用在 slice map channel

	t11 := TypeA{"zhouwude"}
	cTypeA(t11)
	fmt.Println("t11:", t11) //t11: {zhouwude}
	pTypeA(&t11)
	fmt.Println("t11:", t11) //t11: {meibing}
}

// TypeA被拷贝
func cTypeA(t TypeA) {
	t.P1 = "meibing"
}
func pTypeA(t *TypeA) {
	t.P1 = "meibing" //一样会自动解引用
	// (*t).P1 = "没冰"
}
func test2(a int) (x, y int) {
	x = a * 2
	y = a * 2
	defer func() {
		x = 1001
		fmt.Println("____")
	}()
	return
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}
type S1 struct {
	a string
}
