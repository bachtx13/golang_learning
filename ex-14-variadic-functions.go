package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))

	nums := []int{5, 7, 9}
	fmt.Println(sum(nums...))
}
