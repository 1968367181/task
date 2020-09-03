package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("It's the %dst loop\n", i+1)
	}
}
