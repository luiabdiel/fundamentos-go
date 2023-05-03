package main

import (
	"fmt"
	"reflect"
)

func main() {
	num1 := 10.0
	num2 := 20.0

	fmt.Println(num1 + num2)
	fmt.Println(num2 - num1)
	fmt.Println(num1 * num2)
	fmt.Println(num2 / num1)

	fmt.Println(reflect.TypeOf(num1))

	text1 := "texto 1"
	text2 := "texto 2"

	fmt.Println(text1 + text2)
	fmt.Println(reflect.TypeOf(text1))
}
