package main

import "fmt"

func main() {
	texto := "palavra"
	fmt.Println("Quantiade: ", len(texto))

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "l" {
			continue
		}

		if string(texto[i]) == "r" {
			break
		}

		fmt.Println(string(texto[i]))
	}
}
