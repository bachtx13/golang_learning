package main

import "fmt"

func main() {
	fmt.Println(values(2, "abc"))
}
func values(num int, str string) (int, string) {
	return num + 1, str + " concat"
}
