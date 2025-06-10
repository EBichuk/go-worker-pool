package workerpool

import (
	"testing"
	"time"
)

func Test_ClosedWorkerPool(t *testing.T) {
	wp := NewWorkerPool(2)
	wp.Stop()

	time.Sleep(3 * time.Millisecond)

	_, ok := <-wp.closeChan

	if ok {
		t.Errorf("worker pool channel should be closed")
	}

	_, ok = <-wp.jobsQ
	if ok {
		t.Errorf("job channel should be closed")
	}
}

func Test_AddWorker(t *testing.T) {
	wp := NewWorkerPool(2)
	defer wp.Stop()

	wp.AddWorker()
	wp.AddWorker()

	want := 4
	got := wp.countWorkers
	if want != got {
		t.Errorf("Count of workers incorrect: got %v; want %v", got, want)
	}
}

func Test_DeleteWorker(t *testing.T) {
	wp := NewWorkerPool(4)
	defer wp.Stop()
	time.Sleep(3 * time.Millisecond)

	wp.DeleteWorker()
	wp.DeleteWorker()

	want := 2
	got := wp.countWorkers
	if want != got {
		t.Errorf("Count of workers incorrect: got %v; want %v", got, want)
	}

	wp.DeleteWorker()
	wp.DeleteWorker()

	want = 0
	got = wp.countWorkers
	if want != got {
		t.Errorf("Count of workers incorrect: got %v; want %v", got, want)
	}
}
