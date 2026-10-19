package main

import (
	"fmt"
	"slices"
)

func main() {
	countries := []string{"Brazil", "Canada", "Austria", "Lebanon", "Italy", "Japan", "Australia", "South Africa", "New Zealand"}

	fmt.Printf("Slice: %v\nLength: %d\nCapacity: %d\n", countries, len(countries), cap(countries))

	indexToRemove := 3
	countries = append(countries[:indexToRemove], countries[4], countries[5], countries[6], countries[7], countries[8])

	fmt.Printf("Slice: %v\nLength: %d\nCapacity: %d\n", countries, len(countries), cap(countries))

	countries = slices.Delete(countries, 6, 7)
	fmt.Printf("Slice: %v\nLength: %d\nCapacity: %d\n", countries, len(countries), cap(countries))
}
