package main

import "fmt"

func main() {
	resultado := Soma(1, 2)
	fmt.Println(resultado)
}

func Soma(num1 int, num2 int) int {
	resultado := num1 + num2

	return resultado
}
