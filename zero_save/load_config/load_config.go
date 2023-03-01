package main

import (
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	type Map map[string]string
	var m atomic.Value

	m.Store(make(Map))
	var mu sync.Mutex

	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}

	insert := func(key, val string) {
		mu.Lock()
		defer mu.Unlock()
		m1 := m.Load().(Map)
		m2 := make(Map)

		for k, v := range m1 {
			m2[k] = v
		}
		m2[key] = val
		m.Store(m2)
	}

	_, _ = read, insert
}

func S() {
	var config atomic.Value
	config.Store(loadConfig())
	go func() {
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				_, _ = r, c
			}
		}()
	}
}

func loadConfig() int {
	return 1
}

func requests() []string {
	return nil
}
