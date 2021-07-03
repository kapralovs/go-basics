package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { //в данном случае цикл for выступает как цикл while
		sum += sum
	}

	fmt.Println(sum)
}
