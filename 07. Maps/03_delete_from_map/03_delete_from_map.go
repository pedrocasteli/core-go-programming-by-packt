package main

import "fmt"

func main() {
	var populations map[string]int = map[string]int{
		"Sao Paulo":   12000000,
		"Mexico City": 9000000,
		"Sacramento":  538753,
		"Detroit":     649095,
	}

	fmt.Println(populations)

	delete(populations, "Sacramento")

	fmt.Println(populations)
}
