package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// create trace file
	f, err := os.Create("trace.out") // under file
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// start trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
