package main

import "fmt"

func main() {

	switch modulo := 44 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")

	}
	//sin condicion
	value := 90
	switch {
	case value > 100:
		fmt.Println("es mayor a 100")
	case value < 0:
		fmt.Println("es menor que 0")
	default:
		fmt.Println("no condición")

	}

}
