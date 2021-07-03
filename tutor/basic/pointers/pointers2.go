package main

import "fmt"

func main(){
	p:=new(int)
	fmt.Println("Value:",*p)
	*p=8
	fmt.Println("Value:",*p)
}
