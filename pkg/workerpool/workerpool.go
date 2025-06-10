package workerpool

import (
	"fmt"
	"go-worker-pool/pkg/worker"
	"sync"
)

type WorkerPool struct {
	mu           sync.RWMutex
	jobsQ        chan string
	workers      map[int]*worker.Worker
	wg           sync.WaitGroup
	countWorkers int
	closeChan    chan struct{}
}

func NewWorkerPool(countStartWorkers int) *WorkerPool {
	wp := &WorkerPool{
		jobsQ:        make(chan string),
		mu:           sync.RWMutex{},
		wg:           sync.WaitGroup{},
		workers:      make(map[int]*worker.Worker),
		countWorkers: 0,
		closeChan:    make(chan struct{}),
	}

	if countStartWorkers > 0 {
		for range countStartWorkers {
			wp.AddWorker()
		}
	}
	return wp
}

func (wp *WorkerPool) AddWorker() {
	wp.mu.RLock()
	wp.countWorkers++
	fmt.Println("Add worker", wp.countWorkers)
	w := worker.NewWorker(wp.countWorkers, wp.jobsQ)
	wp.workers[wp.countWorkers] = w

	wp.mu.RUnlock()

	wp.wg.Add(1)
	go func() {
		defer wp.wg.Done()
		fmt.Println("Start worker", wp.countWorkers)
		w.Start()
	}()
}

func (wp *WorkerPool) DeleteWorker() {
	wp.mu.RLock()
	if wp.countWorkers < 1 {
		fmt.Println("not have active workers")
		return
	}

	fmt.Println("Delete worker", wp.countWorkers)
	w := wp.workers[wp.countWorkers]

	delete(wp.workers, wp.countWorkers)
	wp.countWorkers--

	wp.mu.RUnlock()

	w.Stop()
}

func (wp *WorkerPool) Stop() {
	close(wp.jobsQ)
	close(wp.closeChan)
	wp.wg.Wait()
}

func (wp *WorkerPool) Submit(job string) {
	wp.jobsQ <- job
}
