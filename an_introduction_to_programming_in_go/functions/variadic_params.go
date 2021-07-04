package main

import "fmt"

func maxNum(nums ...int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	number := maxNum(45, 34, 76)
	fmt.Println(number)
	number = maxNum(23, 9, 2353, 44, 387, 13)
	fmt.Println(number)
}
