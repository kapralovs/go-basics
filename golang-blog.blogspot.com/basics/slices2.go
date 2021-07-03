package main

import "fmt"

func main() {
	names := [4]string{
		"Евгений",
		"Иван",
		"Георгий",
		"сергей",
	}

	fmt.Println(names)

	a := names[0:2]

}
