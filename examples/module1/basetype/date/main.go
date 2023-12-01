package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Printf("%d.%d.%d\n", t.Day(), t.Month(), t.Year()) //21.11.2023
	fmt.Println(time.ANSIC)
	fmt.Println(t.Format("02 Jan 2006 15:04")) //21 Nov 2023 16:31

	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	// Duration 类型表示两个连续时刻所相差的纳秒数
	//单位是纳秒级别1e9 1* 10的九次方
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec转换为纳秒级别
	week_from_now := t.Add(week)
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822))         // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("21 Dec 2011 08:52")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
	// go build tets.go 会生成一个可执行文件 main max是 main: Mach-O 64-bit executable x86_64
}
