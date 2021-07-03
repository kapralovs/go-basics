package main

import "fmt"

type person struct{
        name string
        age int
}

func main(){
	var tom person=person{"Tom",23}
	fmt.Println(tom.name)
}
