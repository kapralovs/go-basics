package main

import "fmt"

func changeValue(x *int){
	*x=(*x)*(*x)
}

func main(){
	d:=5
	fmt.Println("d before:",d)	//значение до изменения 5
	changeValue(&d)	//изменяем значение
	fmt.Println("d after",d)	//значение после изменения 25
}
