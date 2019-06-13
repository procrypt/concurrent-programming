package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan interface{}
	select {
	case <-ch:
		fmt.Println("Never going to unblock..")
	case <-time.After(time.Second*1):
		fmt.Println("Timeout after 1 sec.")
		return
	}
}
