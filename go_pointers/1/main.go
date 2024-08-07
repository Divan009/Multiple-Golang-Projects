package main

import "fmt"

func updateName(x string) {
	x = "x"
	//fmt.Println("x is set %s", x)
}

func main() {
	name := "tifa"

	updateName(name)

	fmt.Println(name)
}
