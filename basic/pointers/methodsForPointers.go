package main

import "fmt"

type person struct{
	name string
	age int
}

func(p *person) updateAge(newAge int){
	(*p).age=newAge
}

func main(){
	var tom person=person{name:"Tom",age:24}
	var tomPointer *person=&tom
	fmt.Println("before:",tom.age)
	tomPointer.updateAge(25)
	fmt.Println("after",tom.age)
}
