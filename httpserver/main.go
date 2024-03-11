package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"net/http/pprof"
	_ "net/http/pprof"

	httpservermetrics "github.com/cncamp/golang/httpserver/metrics"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//cap() capacity 设置slice map channel 的容量可扩展的一个大小
	// fmt.Println(a)
	// v参数 根据这个来设置日志级别
	flag.Set("v", "4")
	flag.Parse()
	//日志级别
	glog.V(2).Info("Starting http server...")
	httpservermetrics.Register()
	// http.HandleFunc("/", rootHandler)
	// c, python, java := true, false, "no!"
	// fmt.Println(c, python, java)
	// err := http.ListenAndServe(":8089", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	// mux.HandleFunc("/metrics", promhttp.Handler())
	mux.Handle("/metrics", promhttp.Handler()) //定义 metrics 方法给普罗米修斯
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(4).Info("entering root handler")
	timer := httpservermetrics.NewTimer()
	// 这里在函数返回的时候执行 记录的是整个函数的执行时间
	defer timer.ObserveTotal()
	user := r.URL.Query().Get("user")
	delay := randInt(10, 2000)
	//毫秒
	time.Sleep(time.Millisecond * time.Duration(delay))
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	fmt.Fprintf(w, "request remote ip %s ", r.RemoteAddr)
	fmt.Fprintln(w, r.Host)
}
func randInt(min, max int) int {
	if min >= max {
		panic("min must be less than max")
	}
	return min + rand.Intn(max-min+1)
}
