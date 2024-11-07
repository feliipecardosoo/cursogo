package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	car1 := 10
	car2 := car1

	fmt.Println(car1, car2)

	car1++

	fmt.Println(car1, car2)
}
