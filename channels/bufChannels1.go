package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 3
	intCh <- 50
	intCh <- 77
	//intCh<-49	блокировка,функция main ждет,когда освободится место в канале

	fmt.Println(<-intCh)
	fmt.Println("The End")
	fmt.Println(cap(intCh))	//емкость канала
	fmt.Println(len(intCh))	//количество элементов в канале
}
