/*
Базовые типы Go следующие


bool логическое выражение true или false (истина или ложь)

string строка

int  int8  int16  int32  int64  целые числа
uint uint8 uint16 uint32 uint64 uintptr

byte  псевдоним для uint8

rune  псевдоним для int32
      представляет Unicode код

float32 float64  числа с плавающей точкой

complex64 complex128  комплексные числа
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
