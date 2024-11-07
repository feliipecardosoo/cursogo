package main

import "fmt"

func main() {
	fmt.Println("Digite seu nome: ")
	var nome string

	fmt.Scanf("%s", &nome)

	fmt.Println("bem vindo," + nome)
	fmt.Println(&nome)
}
