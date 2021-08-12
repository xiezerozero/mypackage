package work

import (
	"fmt"
	"sync"
	"time"
)

// 伪线程池, workerpool是线程池, worker是每一个线程, job是任务, 每个job都从workerpool中获取一个worker处理,处理完毕放回去
type Job struct {
	JobId int
	Msg   string
	Data  interface{}
}

var wg sync.WaitGroup

type Worker struct {
	WorkerId int
	Ch       chan Job
}

func (w Worker) Work(wp *WorkerPool) {
	for {
		wp.WorkerQueue <- w
		select {
		case j := <-w.Ch:
			w.doJob(j, wp)
		}
	}
}

func (w Worker) doJob(j Job, wp *WorkerPool) {
	fmt.Printf("worker id %d start working: job id: %d \n", w.WorkerId, j.JobId)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("worker id %d end working: job id: %d \n", w.WorkerId, j.JobId)
	wp.ResultCh <- j.Msg
	wg.Done()
}

type WorkerPool struct {
	WorkerQueue chan Worker
	JobCh       chan Job
	ResultCh    chan string
}

func NewWorkerPool(workerNum int) *WorkerPool {
	wp := &WorkerPool{
		WorkerQueue: make(chan Worker, workerNum),
		JobCh:       make(chan Job),
		ResultCh:    make(chan string),
	}
	return wp
}

func (wp *WorkerPool) Start() {
	for v := range wp.JobCh { //jobch不阻塞,持续增加job; goroutine数量增加
		go func(v Job) {
			w := <-wp.WorkerQueue //阻塞,直到有可用的worker,
			wg.Add(1)
			w.Ch <- v
		}(v)
	}
	// 同步所有任务完成并且job chan关闭
	wg.Wait()
	close(wp.ResultCh)
}

func App(s1, s2 string) []string {
	s := []string{}
	s = append(s, s1)
	s = append(s, s2)
	return s
}
