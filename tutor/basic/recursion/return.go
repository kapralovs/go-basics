package main

import "fmt"

func main(){
	var a=add(4,5)
	var b=add(20,6)
	fmt.Println(a)
	fmt.Println(b)

	var age,name=addition(4,5,"Tom","Simpson")
	fmt.Println(name)
	fmt.Println(age)
}

func add(x,y int) int{
	return x+y
}

func addition(x,y int,firstName,lastName string) (int,string){
	var z int=x+y
	var fullName=firstName+" "+lastName
	return z,fullName
}
