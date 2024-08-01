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
		attack(evilNinja)
	}
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second)
}
