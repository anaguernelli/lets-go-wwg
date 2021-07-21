package main

import (
	"fmt"
	"time"
)

func main() {

	var birth int
	anoAtual := time.Now().Year()
	fmt.Println("Sua data de aniversário: ")
	fmt.Scan(&birth)
	idade := anoAtual - birth
	fmt.Printf("Voce tem %v anos de idade", idade)

}
