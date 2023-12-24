package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	delete(m, "a")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	_, valueExist := m["a"]
	fmt.Println("valueExist:", valueExist)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
