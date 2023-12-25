package main

import (
	"fmt"
	"math"
)

func uneven(num1, num2 float64) float64 {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}
func Sqrt(x float64) float64 {
	z := x / 2
	z1 := z - (z*z-x)/(2*z)
	for ; uneven(z, z1) > 0.00000000001; z -= (z*z - x) / (2 * z) {
		fmt.Println(uneven(z, z1))
		z1 = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}
