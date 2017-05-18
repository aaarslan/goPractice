package main

import "fmt"

type pokemon struct {
	name string
	attr string
	lvl  int
}

func (p pokemon) lvlName() string {
	return p.name + " " + p.attr
}

func main() {
	p1 := pokemon{"Charizard", "Fire/Flying", 100}
	p2 := pokemon{"Blastoise", "Water", 100}
	p3 := pokemon{"Venusaur", "Grass/Poison", 100}
	fmt.Println(p1.lvlName())
	fmt.Println(p2.lvlName())
	fmt.Println(p3.lvlName())
}
