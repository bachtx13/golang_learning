package main

import "fmt"

func main() {
	nextint := returnAClosure()
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
}
func returnAClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
