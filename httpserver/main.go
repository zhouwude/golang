package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"net/http/pprof"
	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	// v参数 根据这个来设置日志级别
	flag.Set("v", "4")
	flag.Parse()
	//日志级别
	glog.V(2).Info("Starting http server...")
	// http.HandleFunc("/", rootHandler)
	// c, python, java := true, false, "no!"
	// fmt.Println(c, python, java)
	// err := http.ListenAndServe(":8089", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
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
