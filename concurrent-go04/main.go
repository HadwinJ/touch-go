package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok) // ok = FALSE means closed, TRUE means open
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch)
		// ch <- 0
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
