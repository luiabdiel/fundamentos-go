package main

import "fmt"

func main() {
	var listaToda = []int{2, 10, 9, 4, 5, 8, 1, 3}
	var listaMenorQueCinco = make([]int, 0)

	for i := 0; i < len(listaToda); i++ {
		if listaToda[i] < 5 {
			listaMenorQueCinco = append(listaMenorQueCinco, listaToda[i])
		}
	}

	fmt.Println(listaMenorQueCinco)
}
