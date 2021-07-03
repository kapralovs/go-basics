package main

import "fmt"

func main() {
	intCh := make(chan int)

	go func() {
		fmt.Println("Go routine starts")
		intCh <- 5
	}()
	fmt.Println(<-intCh)
	fmt.Println("The End")
}
