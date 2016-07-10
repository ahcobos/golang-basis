package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	names := []string{"leto", "jessica", "paul"}
	checks := make([]bool, 10)
	scores := make([]int, 0, 20)
	sayains := make([]Saiyan, 0, 10)
	sayains2 := []Saiyan{{Name: "ahcobos", Power: 999999}}

	sayains = append(sayains, Saiyan{
		Name:  "andres",
		Power: 10,
	})

	fmt.Println(names, checks, scores, sayains, sayains2)
}
