package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150
	//o valor de 150 está acima da capacidade que um int8 possa suportar (só suporta valor igual ou menor que 127)
	//Portanto, o número da variável será impressa somente a partir do int16 

	fmt.Println(quilometros)
}
