package main

var a1 = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a1)
}

func m() {
	a1 := "O"
	print(a1)
}
