package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)

	evilNinja := "A"

	go channelAttack(evilNinja, smokeSignal)

	fmt.Println("Main getting Completed")
	fmt.Println(<-smokeSignal)
}

func channelAttack(target string, attacked chan bool) {
	select {
	case value := <-attacked:
		fmt.Println("Received value:", value)
	default:
		fmt.Println("No value available in the channel at this moment.")
	}

	fmt.Println("Throwing ninja stars at", target)
	attacked <- true
}
