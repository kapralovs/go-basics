package main

import "fmt"

func manyReturns() (int, bool, string) { //Функция возвращает целое число,логическое выражение и строку.
	return 1, true, "hello"
}

func main() {
	myInt, myBool, myString := manyReturns() //Каждое возвращаемое значение сохраняется в переменной.
	fmt.Println(myInt, myBool, myString)
}
