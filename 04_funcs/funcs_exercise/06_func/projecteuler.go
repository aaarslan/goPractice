package main

import "fmt"

func max(numbers ...int) int {
	var greatest int
	for _, v := range numbers {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}

func main() {
	largest := max(2, 3, 45, 66, 432, 56, 544, 55, 644, 323)
	fmt.Println(largest)
}
