package main

import "fmt"

func main(){
	var a=8
	fmt.Println("a before:",a)
	increment(a)
	fmt.Println("a after:",a)

	add(1,2)
	add(1,2,3)
	add(1,2,3,4)
	add(1,2,3,4,5)
}

func increment(x int){
	fmt.Println("x before:",x)
	x=x+20
	fmt.Println("x after:",x)
}

func add(numbers ...int){
	var sum=0
	for _,number:=range numbers{
		sum+=number
	}
	fmt.Println("sum=",sum)
}
