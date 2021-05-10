package main

import "fmt"

func main() {
	fmt.Println("Start")
	//создание канала и передача в него данных
	fmt.Println(<-createChan(5)) //блокировка
	fmt.Println("End")
}

func createChan(n int) chan int {
	ch := make(chan int) //создаем канал
	ch <- n              //отправляем данные в канал
	return ch
}
