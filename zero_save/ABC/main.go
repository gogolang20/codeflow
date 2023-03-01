package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{1, 4, 7}
	b := []int{2, 5, 8}
	c := []int{3, 6, 9}

	var wg sync.WaitGroup
	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chC := make(chan struct{}, 1)

	wg.Add(9)

	go printA(&wg, a, chA, chB)
	go printB(&wg, b, chB, chC)
	go printC(&wg, c, chC, chA)
	chA <- struct{}{}

	wg.Wait()
}

func printA(wg *sync.WaitGroup, arr []int, chA, chB chan struct{}) {
	for _, val := range arr {
		<-chA
		fmt.Println(val)
		chB <- struct{}{}
		wg.Done()
	}
}

func printB(wg *sync.WaitGroup, arr []int, chB, chC chan struct{}) {
	for _, val := range arr {
		<-chB
		fmt.Println(val)
		chC <- struct{}{}
		wg.Done()
	}
}

func printC(wg *sync.WaitGroup, arr []int, chC, chA chan struct{}) {
	for _, val := range arr {
		<-chC
		fmt.Println(val)
		chA <- struct{}{}
		wg.Done()
	}
}
