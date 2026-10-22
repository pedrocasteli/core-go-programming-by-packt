package main

import "fmt"

func main() {
	var populations map[string]int

	fmt.Println(populations)

	populations = make(map[string]int)

	if populations == nil {
		fmt.Println("Populations map is nil")
	}

	populations["Edinburgh"] = 531370

	fmt.Println(populations)
}
