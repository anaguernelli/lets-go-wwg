package main

import "fmt"

func main() {
	var listaMercado = []string{"azeite", "alho", "cebola"}
	for i := 0; i < 3; i++ {
		fmt.Printf("%d - %s\n", i+1, listaMercado[i])
	}

}
