package main

import "fmt"

func main() {
	var myMap = make(map[int]string)
	myMap[0] = "Hey"
	myMap[1] = "What's"
	myMap[2] = "Up"
	fmt.Println(myMap)
	myMap[2] = "going down"
	fmt.Println(myMap)

	myPokedex := map[int]string{
		1: "Bulbasaur",
		2: "Ivysaur",
		3: "Venusaur",
		4: "Charmander",
		5: "Charmeleon",
		6: "Charizard",
	}
	fmt.Println(myPokedex)

	myTeam := map[int]string{
		1: "Meiko",
		2: "Matsumoto",
		3: "Kamira",
	}
	fmt.Println(myTeam)
	myTeam[3] = "Camilla"
	fmt.Println(myTeam)
	myTeam[4] = "Rias"
	fmt.Println(myTeam)
	delete(myTeam, 4)
	fmt.Println(myTeam)

	var randomSlice = make([]string, 5)
	randomSlice = append(randomSlice,
		"Matsumoto",
		"Meiko",
		"Kamira",
		"Rias",
	)
	fmt.Println(randomSlice)

	for key, val := range myTeam {
		fmt.Println(key, " - ", val)
	}
}
