package work

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestWorker_Work(t *testing.T) {
	workerNum := 100
	wp := NewWorkerPool(workerNum)

	for j := 0; j < workerNum; j++ {
		w := Worker{WorkerId: j, Ch: make(chan Job)}
		go w.Work(wp)
	}
	start := time.Now()
	go wp.Start()

	for i := 0; i < 1000; i++ {
		wp.JobCh <- Job{
			JobId: i,
			Msg:   "jobid" + strconv.Itoa(i),
		}
	}

	close(wp.JobCh)

	for v := range wp.ResultCh {
		fmt.Println(v)
	}
	fmt.Println("time cost: ", time.Since(start))
}
