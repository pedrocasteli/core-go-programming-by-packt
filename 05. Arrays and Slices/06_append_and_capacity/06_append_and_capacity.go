package main

import "fmt"

func main() {
	dogs := []string{"Kira", "Beethoven", "Leia"}

	fmt.Println(cap(dogs))

	dogs = append(dogs, "To be named")

	fmt.Println(len(dogs))
	fmt.Println(cap(dogs))
	fmt.Println(dogs)
}
