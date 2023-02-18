package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("task complete %d\n", id)
	}
}

func main() {
	r := NewRunner(4 * time.Second)

	r.AddTasks(createTask(), createTask(), createTask())

	err := r.Start()
	switch err {
	case ErrInterupt:
		fmt.Println("tasks interrupt")
	case ErrTimeout:
		fmt.Println("time out")
	default:
		fmt.Println("all tasks finished")
	}

}

var (
	ErrTimeout  = errors.New("cannot finish task with the timeout")
	ErrInterupt = errors.New("revived interrupt from OS")
)

// Runner 给定一系列的task 要求再规定的 timeout 里跑完，不然就报错
// 如果系统给了中断信号，也报错
type Runner struct {
	interrupt chan os.Signal
	complete  chan error

	timeout <-chan time.Time
	tasks   []func(int) // task的列表
}

func NewRunner(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t),
		tasks:     make([]func(int), 0),
	}
}

// 添加任务
func (r *Runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterupt
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner) Start() error {
	// replay interrupt from OS
	signal.Notify(r.interrupt, os.Interrupt)

	// run the task
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
