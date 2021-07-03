package main

import "fmt"

func main(){
	hello("Hello Go")
	add(5,8)
	multiply(5,6.5)
}

func hello(helloGo string){
	fmt.Println(helloGo)
}

func add(x,y int){
	var z=x+y
	fmt.Println("x=",x)
	fmt.Println("y=",y)
	fmt.Println("x+y=",z)
}

func multiply(a int,b float32){
	var c=float32(a)*b
	fmt.Println("a*b=",c)
}

