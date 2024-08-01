package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"A", "B", "C", "D"}

	for _, evilNinja := range evilNinjas {
		go concurrentAttack(evilNinja)
	}

	time.Sleep(time.Second * 2)

}

func concurrentAttack(target string) {
	fmt.Println("Throwing ninja stars at", target)
}
