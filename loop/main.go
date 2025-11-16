package main

import "fmt"

func main() {
	// for init(counter); condition; post {} (rumus)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// loop dengan hanya kondisi
	// x := 0
	// for x < 10{
	// 	fmt.Println("Angka : ", x)
	// 	x++
	// }

	// Looping dengan range
	// angka := []int{1,2,3,4,5}
	// for i, v := range angka {
	// 	fmt.Println("index: ", i, "value", v)
	// }

	// loop absen siswa
	siswa := []string{
		"ojak",
		"ahmad",
		"lopson",
		"berax",
		"jhon",
	}
	fmt.Println("Kita absen adik-adik")
	for s := 0; s < len(siswa); s++{
		fmt.Println("Ada", siswa[s])
	}
}