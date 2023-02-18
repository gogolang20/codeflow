package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Golang的并发二 资源池
type DBConnection struct {
	id int32
}

func (D DBConnection) Close() error {
	fmt.Println("database closed # " + fmt.Sprint(D.id))
	return nil
}

var wg sync.WaitGroup
var counter int32

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	return &DBConnection{id: counter}, nil
}

func performQuery(query int, pool *Pool) {
	defer wg.Done()
	resource, err := pool.AcquireResource()
	if err != nil {
		fmt.Println(err)
	}
	defer pool.ReleaseResource(resource)

	t := rand.Int()%10 + 1
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("finish query" + fmt.Sprint(query))
}

func main() {
	p, err := NewPool(Factory, 5)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(10)
	for id := 0; id < 10; id++ {
		go performQuery(id, p)
	}
	wg.Wait()
	p.Close()
}
