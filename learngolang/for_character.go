package main

import "fmt"

func main() {
	var str string = ""
	for i := 0; i < 25; i++ {
		str += "G"
		fmt.Printf("%s\n", str)
	}
}
