package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof" //性能分析的包 非常重要
)

// 性能分析工具
var cpuprofile = flag.String("cpuprofile", "/tmp/cpuprofile", "write cpu profile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		//
		log.Fatal(err)
	}
	// 开始
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var result int
	for i := 0; i < 100000000; i++ {
		result += i
	}
	log.Println("result:", result)
}
