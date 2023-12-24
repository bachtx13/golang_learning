package main

func main() {
	var s []string

	s = make([]string, 3)

	s[0] = "bach"
	s[1] = "yeu"
	s[2] = "ha"
	printArray(s)

	println()

	s = []string{"ha", "yeu", "bach"}
	printArray(s)

	println()

	printArray(s[0:1])

	println()

	printArray(s[1:])

	println()

	printArray(s[:1])
}

func printArray(arr []string) {
	for _, element := range arr {
		print(element + " ")
	}
}
