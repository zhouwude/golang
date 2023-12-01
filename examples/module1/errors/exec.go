// exec.go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func mainexec() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	// 指针
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	// os 包有一个 StartProcess 函数可以调用或启动外部系统命令和二进制可执行文件；
	// 它的第一个参数是要运行的进程，
	// 第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。
	//  这个函数返回被启动进程的 id（pid），或者启动失败返回错误。
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)                                    //直接退出
	}
	fmt.Printf("-- The process id is %v", pid.Pid) //
	// 2nd example: show all processes
	pid, err = os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
	/* Output 1st:
	   The process id is &{2054 0}total 2056
	   -rwxr-xr-x 1 ivo ivo 1157555 2011-07-04 16:48 Mieken_exec
	   -rw-r--r-- 1 ivo ivo    2124 2011-07-04 16:48 Mieken_exec.go
	   -rw-r--r-- 1 ivo ivo   18528 2011-07-04 16:48 Mieken_exec_go_.6
	   -rwxr-xr-x 1 ivo ivo  913920 2011-06-03 16:13 panic.exe
	   -rw-r--r-- 1 ivo ivo     180 2011-04-11 20:39 panic.go
	*/

	// 2) exec.Run //
	/***************/
	// Linux:  OK, but not for ls ?
	// cmd := exec.Command("ls", "-l")  // no error, but doesn't show anything ?
	// cmd := exec.Command("ls")  		// no error, but doesn't show anything ?
	// exec 包中也有同样功能的更简单的结构体和函数
	cmd := exec.Command("ls") // this opens a gedit-window
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	// fmt.Printf("The command is %v", cmd)
	// The command is &{/bin/ls [ls -l] []  <nil> <nil> <nil> 0xf840000210 <nil> true [0xf84000ea50 0xf84000e9f0 0xf84000e9c0] [0xf84000ea50 0xf84000e9f0 0xf84000e9c0] [] [] 0xf8400128c0}
}
