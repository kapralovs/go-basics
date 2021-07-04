package main

import (
	"fmt"
	"time" //необходимо импортировать пакет time, чтобы использовать тип time.Time
)

func main() {
	var now time.Time = time.Now() //метод time.Now возвращает значение time.Time, представляющее текущую дату и время
	var year int = now.Year()      //у значений time.Time имеется метод Year, который возвращает текущий год (или год, выставленный на системных часах вашего компьютера)
	fmt.Println(year)
}
