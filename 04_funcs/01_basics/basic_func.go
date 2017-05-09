package main

import "fmt"

func main() {
	fmt.Println(greeting("Meiko", "Shiraki"))
}

func greeting(fname string, lname string) (string, string) {
	return fmt.Sprint("Hello, ", fname, lname), fmt.Sprint(lname, fname)
}
