package main

import (
	"fmt"
)

type phoneWriter struct{}

func (pw phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}

func main() {
	bytes1 := []byte("+1(234)5678901")
	bytes2 := []byte("+2-345-678-90-00")

	writer := phoneWriter{}
	writer.Write(bytes1)
	writer.Write(bytes2)
}
