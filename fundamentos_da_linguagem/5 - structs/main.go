package main

import "fmt"

func calculosMatematicos(n1, n2 int) (int, int) {
	somar := n1 + n2
	subtrair := n1 - n2
	return somar, subtrair
}

func main() {
	// Tenho que declarar 2 variavies pois estou recebendo 2 retornos kk....
	// se eu não quiser usar uma variavel posso substituir por _
	somaCal, _ := calculosMatematicos(2, 4)
	somaCal2, subCal2 := calculosMatematicos(2, 4)
	fmt.Println(somaCal)
	fmt.Println(somaCal2, subCal2)
}
