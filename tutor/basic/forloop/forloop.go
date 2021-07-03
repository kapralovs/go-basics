package main

import "fmt"

func main(){
	for i:=1;i<10;i++{//внешний цикл
		for j:=1;j<10;j++{//внутренний (вложеный) цикл
			fmt.Print(i*j,"\t")
		}
	}
	fmt.Println()

	var users=[3]string{"Tom","Alice","Kate"}//объявление строкового массива users из 3-х элементов

	for i:=0;i<len(users);i++{	//перебор элементов массива users стандартным видом цикла for
		fmt.Println(users[i])
	}

	for index,value:=range users{	//перебор элементов массива users с выводом индекса элементов и их значений
		fmt.Println(index,value)
	}

	for _,value:=range users{	//перебор элементов массива users без вывода индекса элементов
		fmt.Println(value)
	}

	var numbers=[7]int{-3,-2,-1,0,1,2,3}
	var sum=0
	
	for _,value:=range numbers{
		if value<=0{
			continue	//переходим к следующей итерации
		}
		sum+=value
	}
	fmt.Println("Sum:",sum)	//Вывод значения sum

	var numbers2=[10]int{1,2,3,4,5,6,7,8,9,10}
	var sum2=0

	for _,value:=range numbers2{
		if value>4{
			break	//если число больше 4 выходим из цикла
		}
		sum2+=value
	}
	fmt.Println("Sum2:",sum2)
}
