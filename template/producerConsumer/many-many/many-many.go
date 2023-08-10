package many_many

import (
	"time"

	"codeflow/template/producerConsumer/out"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var tashCH = make(chan Task, 10)
var done = make(chan struct{})

const (
	taskNum int64 = 10000
	// nums    int64 = 100
)

func producer(wo chan<- Task, done chan struct{}) {
	i := int64(0)
	for {
		if i >= taskNum {
			i = 0
		}
		i++
		t := Task{
			ID: i,
		}

		select {
		case wo <- t:
		case <-done:
			out.Println("p exit")
			return
		}
	}
}

func consumer(ro <-chan Task, done chan struct{}) {
	for {
		select {
		case t := <-ro:
			t.run()
		case <-done:
			for t := range ro {
				if t.ID != 0 {
					t.run()
				}
			}
			return
		}
	}
}

func Exec() {
	go producer(tashCH, done)
	go producer(tashCH, done)
	go producer(tashCH, done)
	go producer(tashCH, done)
	go producer(tashCH, done)
	go producer(tashCH, done)

	go consumer(tashCH, done)
	go consumer(tashCH, done)

	time.Sleep(3 * time.Second)
	close(done)
	close(tashCH)
	time.Sleep(3 * time.Second)
}
