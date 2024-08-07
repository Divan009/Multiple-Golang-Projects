package main

import "fmt"

// we are accepting pointer to whatever is stored at tht memory loc
func updateName(x *string) {
	*x = "x"
	fmt.Println("Memory address of x is: ", &x)

	//fmt.Println("x is set %s", x)
}

func main() {
	name := "tifa"

	fmt.Println("Memory address of name is: ", &name)
	//updateName(name)

	//memory address is being stored in m of name
	m := &name
	//fmt.Println("Value of m is: ", m)
	//fmt.Println("Memory address of m is: ", &m)
	//fmt.Println("Value at Memory address which is stored as value in m is: ", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}
