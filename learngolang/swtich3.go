package main

import "fmt"

var k = 6

func main() {
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:

		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
