package main

import "fmt"

func main(){

	//При наличии нескольких вызовов оператора defer вызываться они будут с конца 

	defer finish()	//будет выполняться последним
	defer fmt.Println("Program has been started")	//будет выполняться перед defer finish()
	fmt.Println("Program is working")	//будет выполняться первым
}

func finish(){
	fmt.Println("Program has been finished")
}
