package main

import "fmt"

func main(){
	var quantity int=4	//так же можно записывать объявление и присваивание в сокращенной форме,например так: var quatity=4 или так: quantity:=4
	var length,width float64=1.2,2.4
	var customerName string="Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered",quantity,"sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width,"square meters")
}
