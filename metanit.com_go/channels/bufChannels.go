package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 10
	fmt.Println("Емкость канала:", cap(intCh))
	fmt.Println("Длина канала:", len(intCh))
	intCh <- 3
	fmt.Println("Емкость канала:", cap(intCh))
	fmt.Println("Длина канала:", len(intCh))
	intCh <- 24
	fmt.Println("Емкость канала:", cap(intCh))
	fmt.Println("Длина канала:", len(intCh))
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
}
