package main

import "fmt"

func main() {

	// Forma um de declarar uma variavel

	var nome string = "Felipe Um"
	fmt.Println(nome)

	// Segundo tipo posso não falar o tipo da variavel
	nome2 := "Felipe Dois"
	println(nome2)

	// Posso declarar varias variaveis de uma ves
	var (
		nomePrincipal string = "Felipe"
		segundoNome   string = "Gabriel"
	)
	fmt.Println(nomePrincipal, segundoNome)

	var testeUm, testeDois string = "oi 1", "oi 2"
	fmt.Println(testeUm, testeDois)

	testeTres, testeQuatro := "oi 3", "oi 4"
	fmt.Println(testeTres, testeQuatro)

}
