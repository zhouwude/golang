package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个队列
type Queue struct {
	queue []string
	cond  *sync.Cond
}

// 让 一组goroutine 在满足特定的条件下被唤醒
func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		for {
			q.Enqueue("a")
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		q.Dequeue()
		time.Sleep(time.Second)
	}
}

func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	fmt.Printf("putting %s to queue, notify all\n", item)
	//唤醒通知消费者 有数据了 不唤醒这个q.cond.L会一直锁住
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
