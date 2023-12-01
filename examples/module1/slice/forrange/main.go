package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
)

type func1 func(a int) int

func main() {
	// var a func1
	// q1 := func(a int) int {
	// 	return 10
	// }
	// q1(3)
	mySlice := []int{10, 20, 30, 40, 50}
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
	for index, _ := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 1, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
	im1 := big.NewInt(100)
	im2 := big.NewInt(2)
	im3 := big.NewInt(5)
	im1.Mul(im2, im3)
	im1.Div(im3, im2)
	fmt.Printf("Big Int: %v\n", im1)
}
