package main

import "fmt"

func main() {
	lista := []int{3, 9, 4, 6}
	fmt.Println(lista)

	fmt.Println("Lista pos1 ", lista[0])
	fmt.Println("Lista pos2 ", lista[1])
	fmt.Println("Lista pos3 ", lista[2])
	fmt.Println("Lista pos4 ", lista[3])

	fmt.Println("Tamanho da lista ", len(lista))
}
