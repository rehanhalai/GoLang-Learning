package main

import "fmt"

func main() {
	// var nums []int <- []

	// <- [0,0]
	var nums = make([]int, 2)

	// adds element
	nums = append(nums, 50)

	fmt.Println(cap(nums))

	fmt.Println(nums)
}
