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

	evilNinja := "A"

	go concurrentAttack(evilNinja)
	fmt.Println("Main getting Completed")
	time.Sleep(time.Second * 2)
}

func concurrentAttack(target string) {
	//time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
}
