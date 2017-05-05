package main

import "fmt"

func main() {

	a := 0.98
	b := "Golang is awesome"
	c := 3204
	d := true
	e := "e"
	var myName = "Abdallah"
	var lastName = "Arslan"
	var wifeFirst, wifeLast = "Raya", "Shishani"

	fmt.Printf("%v \t%v \t%v \t%v \t%v \n", a, b, c, d, e)
	fmt.Printf("%T \t%T \t%T \t%T \t%T \n", a, b, c, d, e)
	fmt.Println(myName)
	fmt.Println(lastName)
	fmt.Println(wifeFirst, wifeLast)

}
