package main

import "fmt"

func main() {
	var coffeeSizes [3]string

	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "Small"
	coffeeSizes[1] = "Medium"
	coffeeSizes[2] = "Large"

	fmt.Println(coffeeSizes)
}
