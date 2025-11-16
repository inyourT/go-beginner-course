package main

import "fmt"

func kalkulator(a, b float64, op string) float64{
	var hasil float64

	if op == "+"{
		hasil = a + b
	} else if op == "-"{
		hasil = a - b
	} else if op == "/"{
		hasil = a / b
	}else{
		hasil = a * b
	}

	return  hasil
}

func main() {
	var a,b float64
	var op string

	fmt.Println("===aplkasi kalkulator====") 

	fmt.Print("Masukan angak pertama :")
	fmt.Scanln(&a)

	fmt.Print("Masukan operator :")
	fmt.Scanln(&op)

	fmt.Print("Masukan angka kedua :")
	fmt.Scanln(&b)

	// pangging fungsi
	hasil := kalkulator(a,b,op)
	fmt.Printf("Hasil : %.2f\n", hasil)

}