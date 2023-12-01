package main

import (
	"bytes"
	"fmt"
	"math"
)

/*
但是 Go 语言里有非常灵活的 接口 概念，通过它可以实现很多面向对象的特性。
接口提供了一种方式来 说明 对象的行为：如果谁能搞定这件事，它就可以用在这儿。
跟 swift 中的协议很像 ，遵守协议就是就得实现协议方法而且协议也可以看作一个类型 swift 协议可以有默认实现
类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

实现某个接口的类型（除了实现接口方法外）可以有其他的方法。
一个类型可以实现多个接口。
接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。
即使接口在类型之后才定义，二者处于不同的包中，
被单独编译：只要类型实现了接口中的方法，它就实现了此接口。
*/

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

// 方法大部分都是大写
// 结构体 Square 实现了接口 Shaper 。
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

// 更详细的例子
type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

//其实都有一个一个隐式的类型装换， showValue 切片 直接赋值
func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

type Buffer bytes.Buffer

// interface 嵌套
/*
 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法。

*/
type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}
type Stringer interface {
	String() string
}

func main() {
	sq1 := new(Square) //指针类型
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area()) //The square has area: 25.000000

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	// 定义了一个接口类型的切片引用类型
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	/*
			Shape details:  {5 3}
		Area of this shape is:  15
		Shape details:  &{5}
		Area of this shape is:  25
	*/
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
	/*Value of the asset is 2308.800049
	Value of the asset is 66500.000000*/

	// Is Square the type of areaIntf?
	/*一个接口类型的变量 varI 中可以包含任何类型的值，
	必须有一种方式来检测它的 动态 类型，即运行时在变量中存储的值的实际类型
	interface_type.(type)实际类型
	interface_type.(*type)指针类型
	areaIntf必须是一个接口变量
	*/
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
	/*
		The type of areaIntf is: *main.Square
	areaIntf does not contain a variable of type Circle*/
	// t是转换成功最后返回的转换后的类型
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	// Type Square *main.Square with value &{5}
	// 只需要类型
	switch areaIntf.(type) {
	case *Square:
		// TODO
	case *Circle:
		// TODO
	default:
		// TODO
	}
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
	/*ram #0 is a int
	Param #1 is a float64
	Param #2 is a string
	Param #3 is unknown
	Param #4 is a nil
	Param #5 is a bool*/
	// 上面是判断类型接口也可以当作是一个类型判断接口变量是否是某个接口
	if sv, ok := areaIntf.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	} else {
		// TODO
	}
}

// 可变参数是一个接口的切片
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
