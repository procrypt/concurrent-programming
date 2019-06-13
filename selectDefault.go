package main

import (
	"fmt"
	"time"
)

func main() {
	var c1 chan interface{}
	var c2 chan interface{}

	start := time.Now()
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n", time.Since(start))
	}
	fmt.Println("about to exit..")
}
