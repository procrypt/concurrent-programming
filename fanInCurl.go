package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
)

type data struct {
	Body io.ReadCloser
	Error error
}

func main() {
	fmt.Println("Gs Old:", runtime.NumGoroutine())
	result  := make(chan io.ReadCloser)

	// send
	go send(c1)
	// receive
	go receive(c1, result)

	fmt.Println("Gs New:", runtime.NumGoroutine())
	for i := range result {
		bs, _ := ioutil.ReadAll(i)
		fmt.Println(string(bs))
	}

	fmt.Println("about to exit")
}

func send(c1 chan<- data) {
	url := []string{"www.google.com","www.youtube.com","www.facebook.com","www.amazon.com"}
	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)
	for i:=0;i<gs; i++ {
		for _, v := range url {
			go func() {
				resp, err := http.Get("http://" + v)
				c1 <- data{resp.Body, err}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	close(c1)
}

func receive(c1 <-chan data, result chan<- io.ReadCloser)  {
	for i := range c1 {
		result <- i.Body
	}
	close(result)
}
