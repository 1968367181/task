package main

import "fmt"

func main() {
	var il = 5
	fmt.Printf("An integer: %d, its location memory: %p\n", il, &il)
	var intP *int
	intP = &il
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
