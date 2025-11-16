package main

import "fmt"

func main() {
	// Inisialiasai map
	students := make(map[string]int)

	// menambah data kedalam map
	students["Joko"] = 90
	students["Dodo"] = 97
	students["Ahmad"] = 88
	
	fmt.Println(students)

	// mengecek keberadaan key
	if val, ok := students["Joko"]; ok{
		fmt.Println("Dani score :", val)
	}else{
		fmt.Println("Dani not found")
	}

	// Iterasi
	for name, score := range students{
		fmt.Printf("%s = %d\n", name, score)
	}


}