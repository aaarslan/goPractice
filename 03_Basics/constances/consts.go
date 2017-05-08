package main

import "fmt"

const p string = "Death and Taxes"
const (
	a = iota
	b
	c
	d
)

func main() {
	var q = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
