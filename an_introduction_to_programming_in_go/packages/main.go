package main

import (
	"fmt"
	"serj/an_introduction_to_programming_in_go/packages/math"
	/*
		если мы хотим использовать сокращенное имя пакета, то мы можем использовать псевдоним
		импорт пакета с псевдонимом будет выглядеть следующим образом:
		import m "an_introduction_to_programming_in_go"
		тогда обращение к этому пакету будет таким:
		m.Average()
	*/)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
