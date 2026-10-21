package main

import "fmt"

func main() {
	var population map[string]int = map[string]int{
		"Santiago":         7050000,
		"Melbourne":        5430000,
		"Sydney":           5638830,
		"Tokyo":            14270000,
		"Saint Petersberg": 5600000,
		"Oslo":             728714,
		"Johannesburg":     4803262,
	}

	city := "Nagasaki"
	popul, exists := population[city]

	if exists {
		correctWord := "people"

		if popul == 1 {
			correctWord = "person"
		}

		fmt.Printf("Population of %s is around %d %s.", city, popul, correctWord)
	} else {
		fmt.Printf("%s is not in the map.", city)
	}
}
