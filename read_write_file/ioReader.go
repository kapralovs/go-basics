/*
В качестве io.Reader можно использовать свои кастомные объекты, которые реализуют данный интерфейс. Например:
*/

package main

import (
	"fmt"
	"io"
	"os"
)

type phoneReader string

//В данном случае в качестве интерфейса io.Reader передается объект phoneReader, который считывает цифровые символы из номера телефона.

func (p phoneReader)Read(bs []byte) (int,error){
	count:=0
	for i:=0;i<len(p);i++{
		if(p[i]>='0' && p[i]<='9'){
			bs[count]=p[i]
			count++
		}
	}
	return count,io.EOF
}

func main(){
	phone1:=phoneReader("+1(234)56790-10")
	io.Copy(os.Stdout,phone1)
	fmt.Println()
}
