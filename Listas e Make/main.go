package main

import "fmt"

func main() {
	lista := []string{"Golang", "Java", "C#"}
	lista = append(lista, "Ruby")
	fmt.Println(lista)

	listaNovaa := make([]int, 1)
	listaNovaa[0] = 1
	listaNovaa = append(listaNovaa, 2)
	listaNovaa = append(listaNovaa, 3)

	fmt.Printf("%T\n", listaNovaa)

	somaTotal := 0

	for i := 0; i < len(listaNovaa); i++ {
		somaTotal += listaNovaa[i]
	}

	fmt.Println("Média: ", somaTotal/len(listaNovaa))
}
