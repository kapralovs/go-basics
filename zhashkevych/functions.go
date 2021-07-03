package main

import "fmt"

const pi = 3.1415

func main() {
	printCircleArea()
}

func printCircleArea() {
	circleRadius := 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("Радиус круга %d см \n", circleRadius)
	fmt.Println("Формула для расчета площади круга: A=Pi*r2\n")

	fmt.Printf("Площадь круга: %f32 cм. кв.\n", circleArea)
}
