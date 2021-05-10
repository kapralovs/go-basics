package main

import "fmt"

type person struct{
	name string
	age int}

func main(){
	tom:=person{name:"Tom",age:24}
	var agePointer *int=&tom.age
	*agePointer=35
	fmt.Println(tom.age)
}
