package main

import "fmt"

func main() {
	var numbers []int

	for i := 1; i <= 20; i++ {
		numbers = append(numbers, i)

		if i %3 == 0 && i %5 == 0 {
			fmt.Println("FIZZBUZZ")
		}else if i %3 == 0 {
			fmt.Println("Fizz")
		}else if i %5 == 0 {
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
		}
	}

	// fmt.Println(numbers)
}