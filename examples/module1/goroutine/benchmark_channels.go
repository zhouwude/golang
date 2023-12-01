package main

import (
	"fmt"
	"testing"
)

/*

 标杆分析 Goroutines
如果想排除一部分代码或者更具体的测算时间，
你可以适当使用 testing.B.StopTimer() 和 testing.B.StartTimer()
来关闭或者启动计时器。只有所有测试全部通过，基准测试才会运行。
 标杆分析 Goroutines

*/

func mainbench() {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
	/*sync  4766787	       255.0 ns/op
	buffered 17778139	        72.34 ns/op*/
}

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	//可以自动检测channel 是否关闭
	for range ch {
	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch) //关闭通道
	}()
	for range ch {
	}
}
