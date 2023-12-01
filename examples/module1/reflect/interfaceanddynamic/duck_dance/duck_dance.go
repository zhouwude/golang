package duck

import "fmt"

// 鸭子类型 言简意该 的说法 即忽略对象的真正类型，转而关注对象有没有实现所需 的方法、签名和语义。
type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	// ...
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}
