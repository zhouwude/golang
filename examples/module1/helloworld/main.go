package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
	a, _ := strconv.Atoi("100")
	fmt.Println("a is %n b is \n", a)
	t := strconv.Itoa(100)
	fmt.Printf("-- %s \n", t)
	t1 := strconv.FormatInt(100, 2)
	fmt.Printf("--%s\n", t1)
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")

		for i := 0; i < 5; i++ {
			var v int
			fmt.Printf("%d ", v)
			v = 5
		}
	}
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	for pos, a := range "zhouwude" {
		fmt.Printf("%d---%c", pos, a)
		fmt.Println()
	}

	// this function changes reply:

	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("n:", n)
	if n == *reply {
		fmt.Println("equal")
	} // Multiply: 50
	var u = []int{1, 2, 3}
	fmt.Println(len(u), u[2])

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	where()
	where()
}
func Multiply(a int, b int, reply *int) {
	*reply = a * b
}
