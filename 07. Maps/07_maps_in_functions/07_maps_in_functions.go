package main

import "fmt"

func calcRemainPopul(populations map[string]int, city string, population int) {
	currPopul, exists := populations[city]

	if exists {
		if currPopul >= population {
			populations[city] = populations[city] - population
		}
	} else {
		fmt.Printf("%s is not in the map.", city)
	}
}

func main() {
	var populations map[string]int = map[string]int{
		"Seoul":     9579449,
		"Busan":     3500000,
		"Hiroshima": 1170000,
		"Osaka":     2800000,
	}

	calcRemainPopul(populations, "Oklahoma", 2500)

	fmt.Println(populations)
}
