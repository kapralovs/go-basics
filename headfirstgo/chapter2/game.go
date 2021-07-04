/*
Мини-игра

В чём смысл?

Guess-игра, в которой игрок должен угадать случайное число

1.Генерируется случайное число от 1 до 100
2.Пользователю предлагается угадать это число и сохранить его ответ
3.Если предположение игрока меньше задуманного числа,вывести сообщение (например: "Oops. Your guess was LOW")
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand" //импортируем пакет math/rand
	"os"
	"strconv"
	"strings"
	"time" //импортируем пакет time
)

func main() {
	seconds := time.Now().Unix() //получаем текущую дату и время в формате целого числа
	rand.Seed(seconds)           //функция генератора случайных чисел
	target := rand.Intn(100) + 1 //теперь генерируемые числа будут разные при каждом запуске
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?") //Сообщаем игроку о том,что число выбрано.

	reader := bufio.NewReader(os.Stdin) //Создаём bufio.Reader для чтения ввода данных с клавиатуры.

	success := false	//Объявляем переменную "success" до цикла,для того,чтобы она оставалась в области видимости после цикла
	for guesses := 0; guesses < 10; guesses++ { //Переменная guesses хранит количество сделанных попыток
		fmt.Println("You have", 10-guesses, "guesses left") //Чтобы определить количество оставшихся попыток мы из 10 вычитаем кол-во сделанных попыток
		fmt.Print("Make a guess:")                          //Запросить число.
		input, err := reader.ReadString('\n')               //Прочитать данные,введенные пользователем до нажатия клавиши Enter.

		/*
			Если произошла ошибка,программа выводит сообщение и завершается.
		*/
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //Удаление символа новой строки.
		guess, err := strconv.Atoi(input) //Входная строка преобразуется в число

		/*
			Если произошла ошибка,программа выводит сообщение и завершается.
		*/
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops.Your guess is LOW.") //Если предположение игрока МЕНЬШЕ загаданного числа,то вывести сообщение
		} else if guess > target {
			fmt.Println("Oops.Your guess is HIGH.") //Если предположение игрока БОЛЬШЕ загаданного числа,то вывести сообщение
		} else {
			success = true	//Если игрок угадал число,то выводить сообщение о проигрыше не нужно
			fmt.Println("Good job!You guessed it!") //Если продположение игрока РАВНО загаданному числу,то вывести сообщение
			break                                   //Затем выйти из цикла
		}
	}

	if !success {	//Если игрок не угадал число(если переменная success не равна true)...
		fmt.Println("Sorry,you didn't guess my number.It was:", target)	//...вывести сообщение о проигрыше
	}
}
