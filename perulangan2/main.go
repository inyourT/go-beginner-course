package main

import "fmt"

func main() {
	plat := []string{"A", "B", "C", "D", "E", "A", "A", "B", "B", "AE", "BB", "DK", "DE", "BE"}

	totalA := 0
	totalB := 0
	totalInvalid := 0

	for _, s := range plat {
		switch s{
		case "A":
			totalA += 1
		case "B":
			totalB += 1
		default:
			totalInvalid += 1
		}
	}
	
	fmt.Println("Total yang ada: ")
	fmt.Println("Total A :", totalA)
	fmt.Println("Total B :", totalB)
	fmt.Println("Total No invalid :", totalInvalid)
}