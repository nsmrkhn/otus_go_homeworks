package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrZeroTasks           = errors.New("zero tasks provided")
	ErrZeroWorkers         = errors.New("zero workers provided")
	wg                     = sync.WaitGroup{}
)

type Task func() error

func Run(tasks []Task, n, m int) error {
	var errorsCount int64

	if len(tasks) == 0 {
		return ErrZeroTasks
	}

	if n <= 0 {
		return ErrZeroWorkers
	}

	wg.Add(n)
	channel := produce(tasks)
	for i := 0; i < n; i++ {
		go consume(channel, &errorsCount, m)
	}

	wg.Wait()

	if int(errorsCount) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func produce(tasks []Task) chan Task {
	taskChannel := make(chan Task, len(tasks))
	defer close(taskChannel)
	for _, task := range tasks {
		taskChannel <- task
	}

	return taskChannel
}

func consume(channel chan Task, errorsCount *int64, errMaxCount int) {
	defer wg.Done()
	for task := range channel {
		err := task()
		if err != nil {
			atomic.AddInt64(errorsCount, 1)
		}
		if int(atomic.LoadInt64(errorsCount)) >= errMaxCount {
			return
		}
	}
}
