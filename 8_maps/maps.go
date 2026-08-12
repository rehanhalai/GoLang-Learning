package main

import "fmt"

// maps in GoLang = objects in JS/TS

func main() {

	m := make(map[string]string)
	sm := map[string]int{"price": 10, "areaId": 8}

	fmt.Println(sm)

	fmt.Println(m)

	m["name"] = "GoLang"
	m["area"] = "Backend"

	fmt.Println(m)

	delete(m, "area")

	fmt.Println(m["name"])
	fmt.Println(m)
}
