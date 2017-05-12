package main

import "fmt"

func main() {
	numSlice := make([]int, 1)
	numSlice[0] = 1
	numSlice[0]++
	numSlice = append(numSlice, 5, 4, 32, 54)
	fmt.Println(numSlice)

	numSlice = append(numSlice[:2], numSlice[3:]...)
	fmt.Println(numSlice)
}
