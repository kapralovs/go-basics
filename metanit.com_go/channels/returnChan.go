package main

import "fmt"

func main() {
	fmt.Println("Start")
	//создание канала и получение из него данных
	fmt.Println(<-createChan(5)) //блокировка
	fmt.Println("The End")
}

func createChan(n int) chan int {
	ch := make(chan int) //создаем канал
	ch <- n              //отправляем данные в канал
	return ch            //возвращаем канал
}
