package main

import (
	"fmt"
	workerpool "go-worker-pool/pkg/workerpool"
	"time"
)

func main() {
	p := workerpool.NewWorkerPool(5)

	time.Sleep(1 * time.Second)

	for i := 1; i <= 5; i++ {
		p.Submit(fmt.Sprintf("Job %d", i))
	}

	time.Sleep(1 * time.Second)

	p.DeleteWorker()
	p.DeleteWorker()

	time.Sleep(1 * time.Second)
	for i := 5; i <= 11; i++ {
		p.Submit(fmt.Sprintf("Job %d", i))
	}

	p.Stop()
}
