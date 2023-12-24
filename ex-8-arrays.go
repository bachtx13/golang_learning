package main

import "strconv"

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	for index, element := range arr {
		println(strconv.Itoa(index) + ": " + strconv.Itoa(element))
	}
}
