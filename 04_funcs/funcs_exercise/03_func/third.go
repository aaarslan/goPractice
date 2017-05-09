package main

import "fmt"

func max(numbers ...int) int {
	var greatest int
	for _, x := range numbers {
		if x > greatest {
			greatest = x
		}
	}
	return greatest
}

func main() {
	largest := max(43, 43, 56, 75, 24, 7)
	fmt.Println(largest)
}
