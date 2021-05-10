package main

import "fmt"

func main(){
	var hello string
	hello="Hello world"
	fmt.Println(hello)

	var(
		name string="Red"
		age int=24
		weight float32=75.6
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)

	newVar:="new"

	fmt.Println(newVar)
}
