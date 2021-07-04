package main

import "fmt"

func main() {
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)

	resultString:=fmt.Sprintf("About one-fifth: %0.2f\n",1.0/7.0)
	fmt.Printf(resultString)
}
