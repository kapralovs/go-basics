package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 3
	intCh <- 24
	intCh <- 10
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
}
