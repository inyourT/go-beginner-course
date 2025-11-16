package main

import "fmt"

func main() {
	// array
	og := [...]string{"ana", "topson", "ceb", "jerax", "notail"}

	// slice
	youngPlayer := og[:2]
	oldPlayer := og[2:5]
	coach := append(oldPlayer, "patrick")


	// copy slice
	cop := make([]string, len(coach))
	copy(cop,coach)
	fmt.Println("ini yang copy player: ", cop)

	fmt.Println(og)
	fmt.Println(youngPlayer)
	fmt.Println(oldPlayer)
	fmt.Println(coach)
}