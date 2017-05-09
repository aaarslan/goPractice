package main

import "fmt"

func letItBe(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	m, even := letItBe(2)
	fmt.Println(m, even)
}
