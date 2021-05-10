package main

import "fmt"

func main() {
	var array [3]int = [3]int{1, 2, 3}
	var stringArray [2]string = [2]string{
		"Строка 1",
		"Строка 2"}
	numbers := [5]float32{5.3, 3.7, 1.2, 6.6, 3.5}
	var intArray = [...]int{1, 2, 6, 89, 70, 55, 344}
	fmt.Println(array)
	fmt.Println(stringArray)
	fmt.Println(numbers)
	fmt.Println(intArray)
	colors := [3]string{
		2: "blue", 0: "red", 1: "green"}
}
