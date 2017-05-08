package main

import "fmt"

var a string

func main() {
	fmt.Println("Enter your name")
	fmt.Scan(&a)
	fmt.Printf("Is this your name? %v\n", a)

	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " - FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " - Buzz")
		} else if i%5 == 0 {
			fmt.Println(i, " - Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
