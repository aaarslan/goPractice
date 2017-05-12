package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), rune(i), i%12)
	}
}
