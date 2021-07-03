package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Когда суббота?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Сегодня.")
	case today + 1:
		fmt.Println("Завтра.")
	case today + 2:
		fmt.Println("Через два дня.")
	default:
		fmt.Println("Слишком далеко.")
	}

	//пример switch без условия
	t := time.Now()
	switch { //Switch без условия это тоже самое, что и switch true.
	//Эта конструкция может быть использована как более ясный способ записи длинной цепочки if-then-else.
	case t.Hour() < 12:
		fmt.Println("Доброе утро!")
	case t.Hour() < 17:
		fmt.Println("Добрый день.")
	default:
		fmt.Println("Добрый вечер.")
	}
}
