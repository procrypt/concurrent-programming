package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	done := make(chan struct{})

	// send
	go send(even, odd, done)

	// receive
	receive(even, odd, done)


	fmt.Println("about to exit...")
}

func send(even, odd chan<- int, done chan struct{})  {
	const gs  = 10
	var wg sync.WaitGroup
	wg.Add(gs)
	for i:=0; i<100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	wg.Wait()
	done <- struct{}{}
}

func receive(even, odd <-chan int, done chan struct{}) {
	for {
		select {
		case e := <-even:
			fmt.Println("Even:", e)
		case o := <-odd:
			fmt.Println("Odd: ", o)
		case <-done:
			return
		}
	}
}
