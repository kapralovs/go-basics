package main

import "fmt"

func isEven(n int) bool{
	return n%2==0
}

func isPositive(n int) bool{
	return n>0
}

func sum(numbers []int,criteria func(int) bool) int{
	result:=0
	for _,value:=range numbers{
		if(criteria(value)){
			result+=value
		}
	}
	return result
}

func main(){
	slice:=[]int{-2,4,3,-1,7,-4,23}

	sumOfEvens:=sum(slice,isEven)
	fmt.Println(sumOfEvens)

	sumOfPositives:=sum(slice,isPositive)
	fmt.Println(sumOfPositives)
}
