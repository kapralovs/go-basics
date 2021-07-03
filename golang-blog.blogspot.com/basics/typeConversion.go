package main

import (
	"fmt"
	"math"
)

func main() {
	//Приведение типа
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	//Выведение типа
	v := 42 // переменная v имеет тип int т.к. при объявлении переменной без явного указания типа переменная получает тип значения,указанного после оператора присвоения
	fmt.Printf("v is of type %T\n", v)
}
