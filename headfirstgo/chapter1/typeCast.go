package main

import (
	"fmt"
	"reflect"
)

func main(){
	var myInt int=2	//объявляем переменную с типом int
	fmt.Println(reflect.TypeOf(myInt))	//выводится тип переменной myInt,т.е. int
	fmt.Println(reflect.TypeOf(float64(myInt)))	//выводится тип переменной myInt после приведения к типу float64
}

