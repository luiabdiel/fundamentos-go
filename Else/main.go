package main

import "fmt"

func main() {
	salario := 1500.00

	if salario < 1000.00 {
		salario += (salario * 0.15)
	} else if salario <= 1320 {
		salario += (salario * 0.10)
	} else {
		salario -= (salario * 0.010)
	}

	fmt.Println("Salário final: ", salario)
}
