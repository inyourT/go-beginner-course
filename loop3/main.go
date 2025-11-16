package main

import "fmt"

func main() {
	things := []string{
		"Ransel", "Sepatu", "Hp", "Dompet", "Olive", "Motor",
	}
	fmt.Println("List barang :", things)
	for _, thing := range things{
		fmt.Println("-",thing)
	}
}