package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u usuario
	fmt.Println(u)

	usuario2 := usuario{"Felipe", 21}
	fmt.Println(usuario2)

	usuario3justName := usuario{idade: 20}
	fmt.Println(usuario3justName)
}
