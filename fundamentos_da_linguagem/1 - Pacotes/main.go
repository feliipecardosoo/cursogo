package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("fglibino@gmail.com")
	if erro != nil {
		fmt.Println("Email com o formato invalido")
	}
}
