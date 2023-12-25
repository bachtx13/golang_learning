package main

import "fmt"

func main() {
	nextInt := returnAClosure()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
func returnAClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
