package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)
	changeNumberWithoutPointer(i)
	fmt.Println(i)
	changeNumberWithPointer(&i)
	fmt.Println(i)
}
func changeNumberWithoutPointer(i int) {
	i++
}
func changeNumberWithPointer(i *int) {
	*i++
}
