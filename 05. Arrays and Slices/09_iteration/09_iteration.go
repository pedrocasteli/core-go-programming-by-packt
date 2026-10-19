package main

import "fmt"

func main() {
	artists := []string{"Lady Gaga", "Taylor Swift", "Lana Del Rey", "Madonna", "Beyoncé", "Dua Lipa"}

	for _, artist := range artists {
		fmt.Println(artist)
	}
}
