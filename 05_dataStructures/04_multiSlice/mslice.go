package main

import "fmt"

func main() {

	records := make([][]string, 1)

	student1 := []string{
		"Charmandar",
		"Charizard",
		"100",
	}
	student2 := []string{
		"Bulbasaur",
		"Venusaur",
		"100",
	}
	student3 := []string{
		"Squirttle",
		"Blastoise",
		"100",
	}
	student4 := []string{
		"Pikachu",
		"Raichu",
		"100",
	}
	records = append(records, student1, student2, student3, student4)
	fmt.Println(records)

}
