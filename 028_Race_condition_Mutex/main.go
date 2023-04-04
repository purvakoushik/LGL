package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to the chapter of race conditions in GO")

	var wg = &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var values = []int{}

	wg.Add(3)
	//Some IIFs or lambda or immediately invoked functions as go routines
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("one go R")

		mut.Lock()
		values = append(values, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("second go R")

		mut.Lock()
		values = append(values, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("third go R")

		mut.Lock()
		values = append(values, 3)
		mut.Unlock()

		wg.Done()

	}(wg, mut)

	wg.Wait()
	fmt.Println(values)

}
