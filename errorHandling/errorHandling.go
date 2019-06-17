package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error error
	Response *http.Response

}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)

			for _, url := range urls {
				resp, err := http.Get("http://" + url)
				select {
				case <-done:
					return
				case results <- Result{err, resp}:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"www.google.com", "www.yahoo.com", "badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
