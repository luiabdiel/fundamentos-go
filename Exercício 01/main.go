/* Criar um array de 2 posições de inteiros e armazenar em uma variável a soma total da lista.
A variável deve ser imprimida no console
*/

package main

import "fmt"

func main() {
	ArrayDenumeros := [2]int{2, 3}

	somaTotal := 0

	for i := 0; i < len(ArrayDenumeros); i++ {
		somaTotal += ArrayDenumeros[i]
	}

	fmt.Println(somaTotal)

	// Sem for

	SomaDoArray := ArrayDenumeros[0] + ArrayDenumeros[1]
	fmt.Println(SomaDoArray)
}
