package main

import "fmt"

func main() {
	var population map[string]int = map[string]int{
		"New York":  8580000,
		"Toronto":   3360000,
		"Montreal":  1762949,
		"Boston":    672973,
		"Warsaw":    1860000,
		"Prague":    1409528,
		"Frankfurt": 778589,
	}

	fmt.Println(population)
	fmt.Println(population["Boston"])

	population["Birmingham"] = 1199447

	fmt.Println(population)
}
