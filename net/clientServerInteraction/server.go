package main

import (
	"fmt"
	"net"
)

/*
	объявляем мапу с ключом string, который будет обозначать цвет на английском,
 	и значением string для хранения соответствующего для указанного в ключе названия цвета на русском языке
*/
var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {
	/*
		объявляем объект типа net.Listen,
		который в качестве параметров принимает протокол (tcp),
		 и номаер порта, который он будет слушать (4545)
	*/
	listener, err := net.Listen("tcp", ":4545")

	if err != nil { //если ошибка не равна nil...
		fmt.Println(err) //...то выводим ошибку на экран...
		return           //...и выходим из программы
	}
	defer listener.Close() //закрываем объект listener

	fmt.Println("Server is listening...") //в случае успешного запуска выводим сообщение о том что "сервер слушает"
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn) //запускаем горутину для обработки запроса
	}
}

//обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		//считываем полученные в запросе данные
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0:n]) //на основании полученных данных получаем из словаря перевод
		target, ok := dict[source]
		if ok == false { //если данные не найдены в словаре
			target = "undefined"
		}

		//выводим на консоль сервера диагностическую информацию
		fmt.Println(source, "-", target)
		//отправляем данные клиенту
		conn.Write([]byte(target))
	}
}
