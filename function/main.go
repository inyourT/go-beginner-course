package main

import "fmt"

func hello (a float64, b float64) float64 {
	hasil := a * b
	return hasil
}

func main() {
	var panjang, lebar float64

	fmt.Println("Masukan panjangnya :")
	fmt.Scanln(&panjang)

	fmt.Println("Masukan Lebar :")
	fmt.Scanln(&lebar)

	fmt.Println("Luas persegi :", hello(panjang, lebar), "CM")
}