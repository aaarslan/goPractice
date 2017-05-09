package main

import "fmt"

// func max(numbers ...int) int {
// 	var least int
// 	for _, x := range numbers {
// 		if x > least {
// 			least = x
// 		}
// 		fmt.Println(x, least)
// 	}
// 	return least
// }
//
// func main() {
// 	mini := max(3, 4, 56, 3, 2, 1, 5, 7)
// 	fmt.Println(mini)
// }

func foo(x ...int) {
	fmt.Println(x)
}

func main() {
	foo(4, 5, 6, 32, 4, 5)
}
