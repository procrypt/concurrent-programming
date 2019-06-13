package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j*2
	}
}

func main() {
	jobs := make(chan int, 1000)
	result := make(chan int, 1000)

	for w:=1; w<=10;w++ {
		go worker(w,jobs,result)
	}
	for j:=1; j<=1000; j++ {
		jobs <- j
	}
	close(jobs)
	for a:=1; a<=1000; a++ {
		<-result
	}
}