package main

import "fmt"

func main(){
	users:=[]string{"Kate","Bob","Alex","Steve","Justin","Meg","Robin","Sara"}
//удаляем четвертый элемент
	var n=3
	users=append(users[:n],users[n+1:]...)
	fmt.Println(users)
}
