package main

import "fmt"

func main() {
	const name = "Rohan Joshi"
	const age = 30

	// age := 40
	// name := "test"

	const (
		host = "localhost"
		port = 5000
	)

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(host)
}
