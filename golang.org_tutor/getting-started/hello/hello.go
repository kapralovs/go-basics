package main

import (
	"fmt"

	"golang.org_tutor/getting-started/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
