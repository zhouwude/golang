package person1

type Person struct {
	// 字段未导出
	firstName string
	lastName  string
}

// get 方法
func (p *Person) FirstName() string {
	return p.firstName
}

//  setter 方法
func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
