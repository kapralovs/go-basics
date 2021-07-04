/*
Программа рассчитывает количество краски,необходимое для покраски всех стен
*/
package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) { //Возвращаются 2 значения.Первое-расход краски,второе-сообщает,возникли ли ошибки при выполнении.
	if width < 0 { //Если ширина имеет недопустимое значение,вернуть 0 и признак ошибки.
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 { //Если высота имеет недопустимое значение,вернуть 0 и признак ошибки.
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil //Функция возвращает расход краски и nil если ошибок не возникло.
}

func main() {
	var amount, total float64                  //Объявляем переменные расхода краски для текущей стены,а также для общего расхода по всем стенам.
	amount, err := paintNeeded(4.2, 3.0)       //Вызываем paintNeeded и сохраняем возвращаемое значение.
	fmt.Printf("%.2f liters needed\n", amount) //Выводим расход для первой стены.
	total += amount                            //Прибавляем расход для первой стены к total.
	amount, err = paintNeeded(5.2, 3.5)
	fmt.Printf("%.2f liters needed\n", amount)
	total += amount
	amount, err = paintNeeded(3.4, -2.5)
	fmt.Println(err) //Выводим значение ошибки (или "nil", если ошибки не было)
	fmt.Printf("%0.2f liters needed\n", amount)
	fmt.Printf("Total:%0.2f liters\n", total) //Выводим общий расход краски по всем стенам.
}
