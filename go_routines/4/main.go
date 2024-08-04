package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		//ch <- result{url, err, 0}
		log.Printf("%-20s %s\n", url, err)

	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, err, t}
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}
}

func main() {
	results := make(chan result)

	list := []string{
		"https://amazon.com",
		"https://facebook.com",
		"https://zoho.com",
		"https://meta.com",
		//"https://metdda.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		fmt.Println("I am here waiting to read the channels")
		r := <-results

		fmt.Println("value of r in results channel ", r)

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}

	}

}
