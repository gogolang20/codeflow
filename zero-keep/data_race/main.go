package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Wait    sync.WaitGroup
	Counter int = 0
)

// data race
// go build -race -o main main.go

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Println("Final Counter: ", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(1 * time.Nanosecond) // importance
		value++
		Counter = value
	}
	Wait.Done()
}
