package app

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Runner struct {
	PoolSize  int64
	Jobs      chan bool
	result    uint64
	processor Processor
	wg        sync.WaitGroup
	substring string
}

func NewRunner(poolSize int64, substring string, processor Processor) *Runner {
	jobs := make(chan bool, poolSize)
	return &Runner{
		PoolSize:  poolSize,
		processor: processor,
		Jobs:      jobs,
		substring: substring,
	}
}

func (r *Runner) Run(task string) {
	r.Jobs <- true
	r.wg.Add(1)
	go r.worker(task)
}

func (r *Runner) worker(j string) {
	defer func() {
		r.wg.Done()
		<-r.Jobs
	}()
	text, err := r.processor.Process(j)
	if err != nil {
		fmt.Printf("Error handling %s, error %s\n", j, err.Error())
		return
	}
	count := Count(text, r.substring)
	fmt.Printf("Count for %s: %d\n", j, count)
	atomic.AddUint64(&r.result, uint64(count))
}

func (r *Runner) GetResults() uint64 {
	r.wg.Wait()
	fmt.Printf("Total: %d", r.result)
	return r.result
}
