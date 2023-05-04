package main

import "fmt"

func main() {
	ImprimirMensagem("Calculadora:")

	soma, subtracao, divisao, multiplicacao := Operacoes(1, 2)
	fmt.Println(soma, subtracao, divisao, multiplicacao)
}

func ImprimirMensagem(mensagem string) {
	fmt.Println(mensagem)
}

func Operacoes(num1 int, num2 int) (soma int, subtracao int, divisao int, multiplicacao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	divisao = num1 / num2
	multiplicacao = num1 * num2

	return
}
