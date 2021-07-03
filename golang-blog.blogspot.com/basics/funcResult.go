package main

import "fmt"

//функция swap возвращает две строки
func swap(x, y string) (string, string) {
	return y, x
}

//функция возвращает именованные возвращаемые значения
func namedResults(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(namedResults(17))
}
