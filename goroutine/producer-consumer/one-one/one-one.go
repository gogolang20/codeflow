package one_one

import (
	"code_for_review/goroutine/producer-consumer/out"
	"sync"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var tashCH = make(chan Task, 10)

const (
	taskNum int64 = 10000
)

func producer(wo chan<- Task) {
	for i := int64(1); i <= taskNum; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}
	defer close(wo)
}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(tashCH)
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		consumer(tashCH)
	}(wg)

	wg.Wait()
	out.Println("Exec Done")
}
