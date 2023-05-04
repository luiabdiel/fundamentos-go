/*
	Dado um slice com os itens "2,8,3,10,5,4,7,9,1" que vão de 1 a 10, faça a soma
	em duas variáveis, a primeira com números de 1 a 5 e a segunda com números de
	6 a 10.
	Imprimir os dois resultados no console
*/

package main

import "fmt"

func main() {
	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	somaAteCinco := 0
	somaAteDez := 0

	for i := 0; i < len(lista); i++ {
		if lista[i] <= 5 {
			somaAteCinco += lista[i]
		} else {
			somaAteDez += lista[i]
		}
	}

	fmt.Println(somaAteCinco)
	fmt.Println(somaAteDez)
}
