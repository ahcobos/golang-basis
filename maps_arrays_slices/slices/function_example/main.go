package main

import (
	"fmt"
)

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	saiyans := make([]Saiyan, 0, 10)
	saiyans = append(saiyans, Saiyan{"Goku", 9000})
	fmt.Println(extractPowers(&saiyans))
}

func extractPowers(Saiyans []*Saiyan) []int {
	powers := make([]int, 0, len(Saiyans))
	for _, saiyan := range Saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}
