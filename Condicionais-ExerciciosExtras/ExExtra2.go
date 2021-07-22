//Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.

package main

import (
	"fmt"
)

func main() {
	numero := 10

	switch {
	case numero == 0:
		fmt.Println("Essa variável tem valor 0")
	case numero%2 == 0:
		fmt.Println("O valor dessa variável é múltiplo de 2")
	case numero%3 == 0:
		fmt.Println("O valor dessa variavel é múltiplo de 3")

	default:
		fmt.Println("Número inválido! Não se encaixa como 0, múltiplo de 2 ou 3")

	}
}

//if numero = 0 {
//fmt.Println("Essa variável tem valor 0")
//} else if numero%2 == 0 {
//	fmt.Println("O valor dessa variável é multiplo de 2")
//} else if numero%3 == 0 {
//	fmt.Println("O valor dessa variavel é multiplo de 3")
// } return ??
