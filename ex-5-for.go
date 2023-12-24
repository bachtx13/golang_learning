package main

import "strings"

func main() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	println(strings.Repeat("----------------------------\n", 2))

	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		println(j)
	}
}
