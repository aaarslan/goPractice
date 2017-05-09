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
	largest := max(3, 4, 5, 6, 7, 8, 9, 92)
	fmt.Println(largest)
}
