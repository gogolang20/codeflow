package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 扇入 示例
func generate(message string, interval time.Duration) (chan string, chan struct{}) {
	mc := make(chan string)
	sc := make(chan struct{})
	go func() {
		defer func() {
			close(sc)
		}()
		for {
			select {
			case <-sc:
				return
			default:
				time.Sleep(interval)
				mc <- message
			}
		}
	}()
	return mc, sc
}

func stopGenerate(mc chan string, sc chan struct{}) {
	sc <- struct{}{}
	close(mc)
}

func multipex(mcs ...chan string) (chan string, *sync.WaitGroup) {
	mmc := make(chan string)
	wg := &sync.WaitGroup{}

	for _, mc := range mcs {
		wg.Add(1)
		go func(mc chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			for m := range mc {
				mmc <- m
			}
		}(mc, wg)
	}
	return mmc, wg
}

func main() {
	mc1, sc1 := generate("message from generate 1", 200*time.Millisecond)
	mc2, sc2 := generate("message from generate 2", 200*time.Millisecond)

	mmc, wg1 := multipex(mc1, mc2)

	errs := make(chan error)

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s signal received", <-sc)
	}()

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for m := range mmc {
			fmt.Println(m)
		}
	}()

	// wait for errors
	if err := <-errs; err != nil {
		fmt.Println(err.Error())
	}

	stopGenerate(mc1, sc1)
	stopGenerate(mc2, sc2)

	wg1.Wait()

	close(mmc)
	wg2.Wait()
}
