package main

import (
	"fmt"
	"strings" //необходимо импортировать пакет strings для его дальнейшего использования в программе
)

func main() {
	broken := "G# r#cks!"                     //изначальная "сломаная" строка
	replacer := strings.NewReplacer("#", "o") //возвращает значение strings.Replacer настроенное для замены всех "#" в строке на "о"
	fixed := replacer.Replace(broken)         //вызываем метод Replace для strings.Replacer и передаёт ему строку, в которй должны производиться замены
	fmt.Println(fixed)                        //вывод исправленой строки после замены
}
