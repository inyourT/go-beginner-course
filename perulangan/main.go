package main

import "fmt"

func main() {
	siswa := []string{
		"ojak", "Joko", "Dodo", "Ahmad", "Agung","Dadi", "Golang", "Python", "JavaScript",}

	nilai := []int64{
		80, 88, 90, 98, 78, 65, 34, 50, 45,
	}
	fmt.Println("Nilai yang didapat: ")
	// for id, n := range nilai {
	// 	if(n >= 80) {
	// 		fmt.Println(siswa[id], "Mendapatkan nilai A")
	// 	} else if (n > 66 ){
	// 		fmt.Println(siswa[id], "Mendapatkan nilai B")
	// 	}else{
	// 		fmt.Println(siswa[id], "Mendapatkn nilai E")
	// 	}
	// }

	// BestPractice menggunakan switce
	for i, n := range nilai {
		var grade string 
		switch {
		case n >= 80:
			grade = "A"
		case n > 66:
			grade = "B"
		default:
			grade = "E"
		}
		fmt.Println(siswa[i],"mendapatkan nilai ", grade, n)
	}
}