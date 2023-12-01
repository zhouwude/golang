package main

import (
	"context"
	"fmt"
	"time"
)

//超时，取消操作或者一切异常情况往往需要进行抢占操作或中断后续操作
//context 是设置截止日期同步信号传递请求相关值的结构体
func main() {
	//用于主函数初始化以及测试中 作为顶层的 Context 也及时说一般我们创建 context 都是基于background
	// TODO 在不确定使用什么 context 的时候才会使用
	// viewdeadline 超时时间
	//Withcancel 创建一个可取消的 context
	baseCtx := context.Background()
	//向 context添加键值对
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a")) //
	}(ctx)
	//给 context 设置一个超时时间 获取一个
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel() //
	go func(ctx context.Context) {
		// 在协程周期性的执行一些事情
		ticker := time.NewTicker(1 * time.Second) //间隔一秒写入通道
		//通道取值
		for _ = range ticker.C {
			select {
			//context 是否到了超时时间
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)
	// b
	// enter default
	// child process interrupt...
	// main process exit!
	//看谁先接收到了
	select {
	case <-timeoutCtx.Done():
		time.Sleep(2 * time.Second)
		fmt.Println("main process exit!")
	}
	// time.Sleep(time.Second * 5)
}
