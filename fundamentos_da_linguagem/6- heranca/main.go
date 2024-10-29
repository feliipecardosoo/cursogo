package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudande struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Struct")

	p1 := pessoa{"Felipe", "Cardoso", 21, 170}
	fmt.Println(p1)

	e1 := estudande{p1, "ads", "estacio"}
	fmt.Println(e1)
}
