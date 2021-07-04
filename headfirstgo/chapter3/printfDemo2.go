/*
Демо-вариант программы,показывающей применение функции Printf()
*/
package main

import "fmt"

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")	//Выводим заголовки столбцов
	fmt.Println("----------------------------------")	//Выводим разделительную черту
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paperclips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
}
