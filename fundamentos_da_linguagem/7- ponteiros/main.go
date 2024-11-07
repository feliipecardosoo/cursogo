package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	car1 := 10
	car2 := car1
	var car3 *int = &car1

	fmt.Println(car1, car2, car3)

	car1++

	fmt.Println(car1, car2, *car3)
}
