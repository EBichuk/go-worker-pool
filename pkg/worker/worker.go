package worker

import "fmt"

type Worker struct {
	workerID  int
	closeChan chan struct{}
	jobChan   chan string
}

func NewWorker(workerID int, jobChan chan string) *Worker {
	return &Worker{
		workerID:  workerID,
		closeChan: make(chan struct{}),
		jobChan:   jobChan,
	}
}

func (w *Worker) Start() {
	for {
		select {
		case <-w.closeChan:
			fmt.Printf("Worker %d stopping\n", w.workerID)
			return
		case data, ok := <-w.jobChan:
			if !ok {
				return
			}
			fmt.Printf("Worker %d process %s\n", w.workerID, data)
		}
	}
}

func (w *Worker) Stop() {
	close(w.closeChan)
}
