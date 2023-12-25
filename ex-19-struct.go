package main

import "fmt"

type GirlFriend struct {
	height int
	weight int
}

func main() {
	var hongha = GirlFriend{155, 45}
	var honggha = GirlFriend{155, 45}
	fmt.Println("Same properties:", hongha == honggha)
	fmt.Println("The same person:", &hongha == &honggha)
}
