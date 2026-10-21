package main

import "fmt"

func main() {
	var populations map[string]int = map[string]int{
		"Rio de Janeiro": 6700000,
		"Taipei":         2424159,
		"Naples":         905050,
		"Helsinki":       700000,
		"Auckland":       1547200,
	}

	for key, population := range populations {
		fmt.Printf("%s - %d\n", key, population)
	}
}
