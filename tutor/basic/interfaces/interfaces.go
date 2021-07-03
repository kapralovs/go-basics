package main

import "fmt"

<<<<<<< HEAD
type Vehicle interface{
	move()
}

//структура "Автомобиль"
type Car struct{}

//структура "Самолет"
type Aircraft struct{}


func(c Car) move(){
	fmt.Println("Автомобиль едет")
}

func(a Aircraft) move(){
	fmt.Println("Самолет летит")
}

func main(){
	var tesla Vehicle=Car{}
	var boing Vehicle=Aircraft{}
	tesla.move()
	boing.move()
=======
type Stream interface {
	read() string
	write(string)
	close()
}

func writeToStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	stream.close()
}

//Структура файл
type File struct {
	text string
}

//Структура папка
type Folder struct{}

//Реализация методов для типа *File
func (f *File) read() string {
	return f.text
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

//Реализация методов для типа *Folder
func (f *Folder) close() {
	fmt.Println("Папка закрыта")
}

func main() {
	myFile := &File{}
	myFolder := &Folder{}

	writeToStream(myFile, "hello world")
	closeStream(myFile)
	//closeStream(myFolder)	//Ошибка:тип *Folder не реализует интерфейс Stream
	myFolder.close() //Так можно
}
