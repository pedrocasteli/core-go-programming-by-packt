package main

import "fmt"

func main() {
	var populations map[string]int

	if populations == nil {
	}

	fmt.Println(populations)
	fmt.Printf("Location: %p\n", &populations)

	populations["Casablanca"] = 3200000
	populations["Cairo"] = 10000000

}
