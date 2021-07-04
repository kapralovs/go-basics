package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	fileInfo,err:=os.Stat("my.txt")

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
