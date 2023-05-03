package main

import "fmt"

func main() {
	salario := 3000.00
	modeloDeContrato := "PJ"

	if salario <= 3000.00 && modeloDeContrato == "PJ" {
		salario += (salario * 0.10)
	} else if salario <= 3000.00 && modeloDeContrato == "CLT" {
		salario += (salario * 0.04)
	}

	fmt.Println(salario)
}
