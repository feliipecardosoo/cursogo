package main

import "fmt"

func main() {
	// var nome string
	// fmt.Scanf("%s", &nome)

	// fmt.Println(nome)

	var nomeUm string = "Felipe"
	var nome2 *string
	nome2 = &nomeUm

	fmt.Println(*nome2, nomeUm)

	nomeUm = "Joao"

	fmt.Println(*nome2, nomeUm)
}
