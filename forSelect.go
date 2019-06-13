package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()

	workCounter := 0
	loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		time.Sleep(time.Second*1)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
