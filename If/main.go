package main

import "fmt"

func main() {
	salario := 1320.00
	var bonus float64

	bonus = 0.10

	if salario < 1320 {
		salario += (salario * bonus)
		fmt.Println("Salário com bônus: ", salario)
		return
	}

	fmt.Println("Salário: ", salario)
}
