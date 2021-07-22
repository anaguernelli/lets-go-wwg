//Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.

package main

import (
	"fmt"
)

func main() {
	valor1 := 12
	valor2 := 10
	valor3 := 20

	if valor1 > valor2 && valor1 > valor3 {

		fmt.Printf("O maior valor é %d", valor1)

	} else if valor2 > valor1 && valor2 > valor3 {

		fmt.Printf("O maior valor é %d", valor2)

	} else if valor3 > valor2 && valor3 > valor1 {

		fmt.Printf("O maior valor é %d", valor3)
	}

}
