package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
)

type web struct {
	Body io.ReadCloser
	Err error
}

func main() {
	c1 := make(chan web)
	result := make(chan string)

	fmt.Println("Old:", runtime.NumGoroutine())
	go send(c1)

	go receive(c1, result)

	for i := range result {
		fmt.Println(i)
	}

	fmt.Println("About to exit..")

}

func send(c chan<- web)  {
	var wg sync.WaitGroup
	urls := []string{"www.google.com", "www.amazon.com","www.facebook.com","www.yahoo.com"}
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, err :=  http.Get("http://"+u)
			c <- web{resp.Body, err}
		}(url)
	}
	wg.Wait()
	close(c)
}

func receive(c <-chan web, r chan string)  {
	for i := range c {
		bs, _ := ioutil.ReadAll(i.Body)
		r <- string(bs)
	}
	close(r)
}