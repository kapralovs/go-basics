package main

import "fmt"

func main() {
	employee := struct {
		name   string
		sex    string
		age    int
		salary int
	}{
		name:   "Вася",
		sex:    "М",
		age:    25,
		salary: 1500,
	}

	fmt.Printf("%+v\n", employee)
}
