//Considerando os times do item anterior, crie uma slice para representar cada um.
//1 Adicione o jogador Luis Inácio no time vermelho usando a função append()
//2 Printe na tela os nomes dos jogadores do time vermelho 
package main
impor("fmt")
func main(){

    timeAmarelo := []string{"Fernando", "Joao", "Lucia", "Mariana", "Ana"}
	timeVermelho := []string{"Helena", "Jonas", "Jose", "Juliana"}

	fmt.Println("O time amarelo é formado por: ", timeAmarelo)
	timeVermelho = append(timeVermelho, "Luis Inacio")
	fmt.Println("O time vermelho é formado por: ", timeVermelho)

	
}