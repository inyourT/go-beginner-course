package main

import "fmt"

func main() {
	siswa := map[int32]string{
		1: "Ahmad",
		2: "Akbar",
		3: "Joko",
	}

	siswa[4] = "Dodo"

	fmt.Println(siswa)

	country := make(map[string]string)

	country["ID"] = "Indonesia"
	country["JP"] = "Jepang"
	
	fmt.Println(country)

}
