package main

import (
	"fmt"
	"sync"
)

var counter int = 0 //общий ресурс

func main() {
	ch := make(chan bool) //канал
	var mutex sync.Mutex  //определям мьютекс
	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The End")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() //блокируем доступ к переменной counter
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Gorutine", number, "-", counter)
	}
	mutex.Unlock() //деблокируем доступ
	ch <- true
}
