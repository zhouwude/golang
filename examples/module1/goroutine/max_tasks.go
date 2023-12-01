package main

import "fmt"

// . 限制并发数
// 程序中的协程通过使用带缓冲通道（这个通道作为一个 semaphore 被使用）
// 来调整资源的使用，实现了对内存等有限资源的优化。
const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	// do something
}

/********不超过 MAXREQS 的请求将被处理并且是同时处理，因为当通道 sem 的缓冲区全被占用时，
函数 handle 被阻塞，直到缓冲区中的请求被执行完成并且从 sem 中删除之前，不能执行其他的请求
这里同时可以创建50个任务当到达51时候由于缓冲区满了阻塞了知道接收者接受值了缓冲区有位置了才继续执行任务
*/
// sem 就像一个 semaphore （信号量），表示一个在一定条件的程序中的一个标志变量的技术术语
func handle(r *Request) {
	fmt.Println("-----------------------")
	//信号量设置 50个任务带缓冲 不阻塞 满了之后阻塞无法进行任务
	sem <- 1 // 不管我们往里面放什么
	process(r)
	//缓冲区有固定上线 这里就是释放一个位置然后等待的任务就可以执行----
	<-sem //释放一个位置 缓冲区中有一个空位置:可以开始下一个请求
}

func server(service chan *Request) {
	for {
		request := <-service
		go handle(request)
	}
}

func maintask() {
	service := make(chan *Request)
	go server(service)
}
