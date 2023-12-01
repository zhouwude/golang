package main

type func1 func(int, int)

func main() {
	// var a *int
	// *a += 1
	// DoOperation(1, increase)
	DoOperation(1, decrease)
}

// ab都是 int
func increase(a, b int) int {
	return a + b
}

func DoOperation(y int, f func1) {
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is:", a-b)
}
