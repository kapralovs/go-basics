package main

import "fmt"

func main() {
	ages := map[string]int{
		"Максим": 20,
		"Олег":   24,
		"Саня":   28,
	}

	age, exists := ages["Антон"]
	if !exists {
		fmt.Println("Антона нет в списке")
	} else {
		fmt.Printf("Антону %d лет\n", age)
	}
}
