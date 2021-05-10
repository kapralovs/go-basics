package main

import "fmt"

type Car struct{}
type Aircraft struct {}

func(c Car) move(){
	fmt.Println("Автомобиль едет")
}

func(a Aircraft) move(){
	fmt.Println("Самолет летит")
}

func driveCar(c Car){
	c.move()
}

func driveAircraft(a Aircraft){
	a.move()
}

func main(){
	var tesla Car=Car{}
	var boing Aircraft=Aircraft{}
	driveCar(tesla)
	driveAircraft(boing)
}
