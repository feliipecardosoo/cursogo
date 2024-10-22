package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	var u usuario
	fmt.Println(u)
}
