package main

import (
	"fmt"
	"reflect"
)

func main() {
	const taxa = 10.0 // Não é possível atribuir uma variável com :=

	fmt.Println(taxa)                 // 10
	fmt.Println(reflect.TypeOf(taxa)) // float64
}
