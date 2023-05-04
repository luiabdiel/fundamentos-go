package main

import "fmt"

func main() {
	var listaToda = []int{2, 10, 9, 4, 5, 8, 1, 3, 2, 11}

	segundaLista := listaToda[:3]
	terceiraLista := listaToda[4:]
	ultimoItem := listaToda[len(listaToda)-1:]

	fmt.Println(segundaLista)  // [2 10 9]
	fmt.Println(terceiraLista) // [5 8 1 3 2 11]
	fmt.Println(ultimoItem)    // [11]
}
