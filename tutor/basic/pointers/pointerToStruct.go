package main

import "fmt"

type person struct{
	name string
	age int
}

func main(){
	tom:=person{name:"Tom",age:13}
	var tomPointer *person=&tom		//создаем указатель tomPointer на адрес экземпляра структуры tom
	tomPointer.age=29
	fmt.Println(tom.age)
	(*tomPointer).age=32	//меняя значение указателя мы меняем значение в ячейке памяти на которую он указывает,т.е. значение поля age
	fmt.Println(tom.age)
}
