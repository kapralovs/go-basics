package main

import "fmt"

type contact struct{
	email string
	phone string
}

type person struct{
	name string
	age int
	contactInfo contact
}

func main(){
	var tom=person{
		name:"Tom",
		age:24,
		contactInfo:contact{
			email:"tom@mail.ru",
			phone:"+1234567890"},
	}
	tom.contactInfo.email="supertom@mail.ru"

	fmt.Println(tom.contactInfo.email)
	fmt.Println(tom.contactInfo.phone)
}
