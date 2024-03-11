package main

import (
	"fmt"
	// 文件名和 package 不一样的时候 go.mod  module
	format "github.com/cncamp/golang/examples/module1/package_tes/do-format"
	"github.com/cncamp/golang/examples/module1/package_tes/math"
	"github.com/cncamp/golang/examples/module1/package_tes/pack1"
	"github.com/cncamp/golang/examples/module1/package_tes/trans"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	fmt.Printf("Float from package1: %f\n", pack1.Pack1Float)

	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
	fmt.Println(trans.Pi) //3.141592653589793
	// histogram_quantile(0.95, sum(rate(httpserver_execution_latency_seconds_bucket[5m])) by (le))

}
