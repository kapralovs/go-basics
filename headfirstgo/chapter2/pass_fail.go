//pass_fail сообщает, сдал ли пользователь экзамен
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ") //Запрашиваем значение у пользователя

	/*
	   В переменную reader помещается объект NewReader из пакета bufio принимающий в качестве аргумента стандартный поток вывода os.Stdin
	*/

	reader := bufio.NewReader(os.Stdin) //Создать "буферизированное средство чтения" для получения текста с клавиатуры

	/*
		В переменную input помещается пепеменная reader, вызывающая метод ReadString
	*/
	input, err := reader.ReadString('\n') //Возвращает текст,введенный пользователем до нажатия клавиши "Enter"
	//Возвращенный код ошибки снова сохраняется в переменной err
	if err != nil { //Если переменная err не пуста,тогда выполняется действие между фигурных скобок

		log.Fatal(err) //Сообщаем об ошибке и завершаем программу
	}

	input = strings.TrimSpace(input)            //удаляем символ новой строки из входной строки
	grade, err := strconv.ParseFloat(input, 64) //Преобразовываем строку в значение float64

	var status string //Переменная, в которую будет записываться результат экзамена

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}

	fmt.Println("A grade of", grade, "is", status)	//Выводим введенное значение и результат сдачи экзамена
}
