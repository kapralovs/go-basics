/*
Поскольку запись в стандартный потое os.Stdout-довольно распространенная задача,то вместо функций Fprint/Fprintln/Fprintf применяются их двойники:Println(),Print(),Printf() соответственно,которые по умолчанию выводят данные в os.Stdout:
*/

package main

import "fmt"

type person struct{
	name string
	age int32
	weight float64
}

func main(){
	tom:=person{
		name:"Tom",
		age:24,
		weight:68.5}

	fmt.Printf("%-10s %-10d %-10.3f\n",
		tom.name,tom.age,tom.weight)
	fmt.Print("Hello ")
	fmt.Println("cold!")
}
