package main

import "fmt"

func main() {
	ImprimirMensagem("Hello, Go")
}

func ImprimirMensagem(mensagem string) {
	mensagem += ". Fine?"

	fmt.Println(mensagem)
}
