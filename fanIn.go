package main

import (
	"fmt"
	"sync"
)

func main() {
	even  := make(chan int)
	odd   := make(chan int)
	fanIn := make(chan int)

	// send
	go func(even, odd  chan<- int) {
		for i:=0; i<100; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(even)
		close(odd)
	}(even,odd)

	// receive
	go receive(even, odd, fanIn)

	for i := range fanIn {
		fmt.Println(i)
	}

	fmt.Println("About to exit...")
}

func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(even <-chan int) {
		for i := range even {
			fanIn <- i
		}
		wg.Done()
	}(even)

	go func(odd <-chan int) {
		for i := range odd {
			fanIn <- i
		}
		wg.Done()
	}(odd)
	wg.Wait()
	close(fanIn)
}