package main

import (
	"fmt"
	"time"
)

/*
Go资源池简版
https://blog.csdn.net/finghting321/article/details/106492915/
*/

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}
	return &t
}

func (t *Task) Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
	worker_num   int
	JobsChannel  chan *Task
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		worker_num:   cap,
		JobsChannel:  make(chan *Task),
	}
}

func (p *Pool) worker(work_ID int) {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("worker ID ", work_ID)
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.worker_num; i++ {
		fmt.Println("start worker:", i)
		go p.worker(i)
	}
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
	close(p.JobsChannel)
	fmt.Println("close JobsChannel")

	close(p.EntryChannel)
	fmt.Println("close EntryChannel")
}

func main() {
	t := NewTask(func() error {
		fmt.Println("create task: ", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})

	p := NewPool(3)

	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	p.Run()
}
