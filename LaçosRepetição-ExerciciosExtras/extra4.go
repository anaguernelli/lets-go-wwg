package main

import (
	"fmt"
	"os"
)

func main() {
	var v int
	fmt.Printf("Insira um número:\n")
	fmt.Fscanf(os.Stdin, "%d", &v)
	for (v % 2) != 0 {
		fmt.Printf("Insira um número:\n")
		fmt.Fscanf(os.Stdin, "%d", &v)
	}
	fmt.Printf("Agora sim podemos dividir igualmente entre mim e você!")

}

//Fscanf   ->   F de file, lê do arquivo, mas basicamente o mesmo
//que na função scanf
//and os.stdin is a file!!
