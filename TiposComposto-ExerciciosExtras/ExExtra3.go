//Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.
//Passos:
//Criar um mapa com chave do tipo string e valor do tipo string (map[string]string)
//onde a chave seja o nome da cidade e o valor o nome do país.
//Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país.
//Esse segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor a quantidade de menções.
//Printe na tela os valores do último mapa.
package main

import "fmt"

func main() {

	mapa1 := map[string]string{
		"Rio de Janeiro": "Brasil",
		"Sao Paulo":      "Brasil",
		"Roma":           "Italia",
	}

	mapaFrequencia := make(map[string]int)

	for _, valor := range mapa1 {
		mapaFrequencia[valor] += 1
	}

	fmt.Println(mapaFrequencia)

}
