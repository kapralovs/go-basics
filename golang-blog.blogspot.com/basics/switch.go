/*
Использование switch в Go во многом схоже с использованием for и if.

Условное выражение не должно быть заключено в круглые скобки ( ),
но фигурные скобки { } обязательны.

Также выражение switch может начинаться с инструкции,
которая будет выполнена перед проверкой условия.
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go выполняется на ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS.X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//freebsd,openbsd,
		//windows
		fmt.Printf("%s", os)
	}
}
