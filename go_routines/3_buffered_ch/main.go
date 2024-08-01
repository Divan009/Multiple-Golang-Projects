package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)
	fmt.Println("before starting channel")
	go func() {
		fmt.Println("in channel")

		time.Sleep(time.Second * 3)
		channel <- "First Msg"
	}()

	fmt.Println("after starting channel")

	fmt.Println(<-channel)
}
