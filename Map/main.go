package main

import "fmt"

func main() {
	// m := map[string]int{"sp": 12000000, "ba": 15000000}
	// fmt.Println(m)

	m := make(map[string]int)

	m["sp"] = 12000000
	m["ba"] = 15000000
	m["pr"] = 11000000

	valor, foiEncontrado := m["rj"]
	if foiEncontrado {
		fmt.Println(valor)
	} else {
		fmt.Println("Essa chave não existe")
	}

	fmt.Println(m)
}
