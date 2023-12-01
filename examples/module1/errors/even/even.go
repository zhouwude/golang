package even

func Even(i int) bool { // Exported function
	return i%2 == 0 //奇数
}

func Odd(i int) bool { // Exported function
	return i%2 != 0 //偶数
}
