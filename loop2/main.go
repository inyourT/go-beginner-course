package main

import "fmt"

func main() {
	clubs := []string{
		"Borneo", "Persija", "Malut", "Persib", "Psim Yogakarta", "Bhayangkara",
	}
	fmt.Println("Klasemen Super leag")
	for index, club := range clubs {
		// fmt.Printf("Peringkat %v adalah %v\n",
		// index+1, club)
		fmt.Println("Peringat", index+1, "adalah", club)
	}
}