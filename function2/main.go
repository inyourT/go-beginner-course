package main

import "fmt"

func heroDescription(name string, skills []string) {
	fmt.Println("Nama hero", name)
	for i, v := range skills{
		fmt.Printf("Skill %d: %s\n", i+1, v)
	}
}

func main() {
	hero := "Riki"
	skillHero := []string{"Smoke screen", "Blink strike", "Trick of the trade", "Cloak and dagger"}

	heroDescription(hero, skillHero)
}