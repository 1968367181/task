package main

import "fmt"

func main() {
	fmt.Printf(Season(4))
}
func Season(month int) string {
	switch month {
	case 1, 2, 3:
		return "spring"
	case 4, 5, 6:
		return "summer"
	case 7, 8, 9:
		return "autumn"
	case 10, 11, 12:
		return "winter"
	default:
		return "don't exist"
	}
}
