package main

import "fmt"

func main(){
	var people=map[string]int{"Tom":1,"Bob":2,"Sam":4,"Alice":8}
	fmt.Println(people)
	people["Bob"]=32
	fmt.Println(people["Bob"])

	if val,ok:=people["Tom"]; ok{
		fmt.Println(val)
	}

	for key,value:=range people{
		fmt.Println(key,value)
	}

	users:=make(map[string]int)	//альтернативный вариант создания хеш-таблицы

	users["Kate"]=153
	fmt.Println(users)

	delete(people,"Tom")
	fmt.Println(people)
}
