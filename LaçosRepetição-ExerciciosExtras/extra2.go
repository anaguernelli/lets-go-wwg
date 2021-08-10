package main

import (
	"fmt"
)

type Apt struct {
	numero       int
	proprietaria string
	vagaGaragem  bool
}

func main() {
	ap1 := Apt{
		numero:       102,
		proprietaria: "Ana",
		vagaGaragem:  false,
	}

	ap2 := Apt{
		numero:       203,
		proprietaria: "Gabi",
		vagaGaragem:  true,
	}

	ap3 := Apt{
		numero:       204,
		proprietaria: "Dani",
		vagaGaragem:  false,
	}
	apartamentos := []Apt{ap1, ap2, ap3}
	for _, Apt := range apartamentos {
		possuiVaga := "Nao"
		if Apt.vagaGaragem {
			possuiVaga = "Sim"
		}

		fmt.Printf("O apartamento %d tem como proprietária %s. Tem vaga de garagem? %s\n", Apt.numero, Apt.proprietaria, possuiVaga)

	}

}
