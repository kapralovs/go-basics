package main

import "fmt"

func main(){
	defer finish()	//будет выполняться в конце программы
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

func finish(){
	fmt.Println("Program has been finished")
}
