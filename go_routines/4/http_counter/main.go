package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> You got %d </h1>", <-nextID)
}

func counter() {
	for i := 0; ; i++ {
		nextID <- i
		time.Sleep(time.Second)
	}
}

func main() {
	go counter()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
