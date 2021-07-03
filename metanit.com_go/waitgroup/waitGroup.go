package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполение \n", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
	fmt.Println("Горутины завершили выполнение")
}
