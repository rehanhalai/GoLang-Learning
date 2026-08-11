package main

import "fmt"

func main() {

	// default values when intialized
	// int -> 0
	// string -> ""
	// bool -> false

	var names [3]string

	names[0] = "ajay"
	fmt.Println(names[0])
	fmt.Println(names)

	// array length
	// fmt.Println(len(names))

	nums := [3]int{1, 2, 3}
	fmt.Println(nums)

}
