package main

import "fmt"

func main() {
	arr := []string{"bach", "yeu", "ha"}
	for _, element := range arr {
		fmt.Printf("%s ", element)
	}
	fmt.Println()
	obj := map[string]string{"a": "apple", "b": "banana", "c": "coconut"}
	for k, v := range obj {
		fmt.Printf("%s: %s\n", k, v)
	}
}
