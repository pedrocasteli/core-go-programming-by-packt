package main

import "fmt"

func main() {
	var countries [6]string = [6]string{"Belgium", "Australia", "Japan", "New Zealand", "Germany", "Thailand"}

	var asianCountries []string = countries[2:6]

	fmt.Println(asianCountries)
}
