package main

import "fmt"

// function which can receive parameters upto n count

func sum(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

func main() {
	// fmt.Println(sum(2, 3, 4, 6, 5, 2, 3, 2, 6, 3, 6, 6, 3, 2))

	nums := []int{1, 5, 2, 5, 2, 6, 32, 4, 20}
	fmt.Println(sum(nums...))
}
