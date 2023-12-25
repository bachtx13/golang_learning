package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("cauuu %d\n", i)
	}
	fmt.Println("yeuuuu")
}
