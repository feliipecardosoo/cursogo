package main

import "fmt"

func somar(n1 uint8, n2 uint8) uint8 {
	return n1 + n2
}

var f = func(txt string) {
	fmt.Println(txt)
}

func calculosMatematicos(n1, n2 int) (int, int) {
	somar := n1 + n2
	subtrair := n1 - n2
	return somar, subtrair
}

func main() {
	soma := somar(3, 5)
	fmt.Println(soma)

	f("testando")

	somaCal, subCal := calculosMatematicos(2, 4)
	fmt.Println(somaCal, subCal)
}
