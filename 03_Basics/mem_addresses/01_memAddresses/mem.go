package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	a := 42
	var meters float64
	fmt.Println("a - ", a)
	fmt.Println("a is located in ", &a)
	fmt.Printf("%d\n", &a)
	fmt.Println("How many meters?")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println("Meters converted to yards is ", yards)
}
