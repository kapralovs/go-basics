package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 5
	close(intCh)
	//intCh <- 24	//ошибка - канал уже закрыт
	fmt.Println(<-intCh) //10
	fmt.Println(<-intCh) //3
	fmt.Println(<-intCh) //0
}
