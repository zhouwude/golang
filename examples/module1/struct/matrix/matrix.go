package matrix

// 小写不让外部包直接调用
// 当结构体的命名以大写字母开头时，该结构体在包外可见。这里不可见
type matrix struct {
}

// 如何强制使用工厂方法
// 通常情况下，为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体
func NewMatrix(params int) *matrix {
	m := new(matrix) // 初始化 m 指针
	return m
}
