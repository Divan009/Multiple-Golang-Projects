package main

import "fmt"

// *pointer is accessing the value of memory address stored in var pointer
func main() {
	var a int
	a = 32
	fmt.Println("Value stored in a: ", a)
	fmt.Println("Memory address where a is stored: ", &a)
	var pointer *int
	fmt.Println("After initialization of pointer: ", pointer)
	pointer = &a
	fmt.Println("After &a = pointer, value at pointer: ", pointer)
	fmt.Println("Memory address of pointer: ", &pointer)
	fmt.Println("*pointer is accessing the value of memory address stored in var pointer: ", *pointer)
}
