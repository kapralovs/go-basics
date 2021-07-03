package main

import "fmt"

type library []string

func (l library) print(){
	for _,value:=range l{
		fmt.Println(value)
	}
}

func main(){
	var lib library=library{"Book 1","Book 2","Book 3"}
	lib.print()
}
