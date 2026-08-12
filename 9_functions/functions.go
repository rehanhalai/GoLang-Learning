package main

import "fmt"

// func add(a int, b int) int {
// func add(a, b int) int {
// return a + b
// }

// func getLangs() (string, string, bool) {
// 	return "html", "css", true
// }

func processUser(FN func(a, b string) string) string {
	return FN("john", "doe")
}

func main() {
	// fmt.Println(add(10, 20))

	// l1, l2, b1 := getLangs()

	// fmt.Println(l1, l2, b1)

	// fmt.Println(getLangs())

	FN := func(a, b string) string {
		return a + b
	}

	fmt.Println(processUser(FN))

}
