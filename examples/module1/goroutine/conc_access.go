package main

import (
	"fmt"
	"time"

	"strconv"
)

// 这是一个简化的例子，并且它不应该在这种情况下应用，但是它展示了如何在更复杂的情况下解决问题。
type Person struct {
	Name string

	salary float64

	chF chan func()
}

func NewPerson(name string, salary float64) *Person {
	// chan func() 函数类型的通道无参数 无返回值
	p := &Person{name, salary, make(chan func())}

	go p.backend()

	return p

}

func (p *Person) backend() {
	//for range 通道
	// 相当于chF通道指定接受者 这里备阻塞 等待发送者
	// 使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：
	// 通道没有值的话 会阻塞 知道可用的发送者
	for f := range p.chF {

		f()

	}

}

// 设置 salary.

func (p *Person) SetSalary(sal float64) {
	// 加入函数
	p.chF <- func() { p.salary = sal }

}

// 取回 salary.

func (p *Person) Salary() float64 {

	fChan := make(chan float64)
	//激活 for range
	p.chF <- func() {
		fmt.Println("Salary point 2")
		time.Sleep(3e9)   //这里 模拟三秒钟 三秒过后才有可用的发送者函数才返回String()输出内容
		fChan <- p.salary //发送者
		fmt.Println("Salary point 3")
	}
	fmt.Println("Salary point 1")
	// 懂了 这里如果fChan没有可用的发送者阻塞一直不返回 fChan <- p.salary 执行之后 Salary()返回
	return <-fChan //获取通道的值也就是消费者如果没有可用接受者 则阻塞
	/*Salary point 1
	Salary point 2
	Salary point 3
	Person - name is: Smith Bill - salary is: 2500.50*/
}

//自定义格式化输出
func (p *Person) String() string {

	return "Person - name is: " + p.Name + " - salary is: " +

		strconv.FormatFloat(p.Salary(), 'f', 2, 64)

}

func mainconcu() {

	bs := NewPerson("Smith Bill", 2500.5)
	// 调用 Strings 函数
	fmt.Println(bs)

	// bs.SetSalary(4000.25)

	// fmt.Println("Salary changed:")

	// fmt.Println(bs)
	/*Person - name is: Smith Bill - salary is: 2500.50
	Salary changed:
	Person - name is: Smith Bill - salary is: 4000.25*/
}
