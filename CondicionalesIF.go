package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2
	pass := 123
	email := "correo"

	if valor1 == 1 {
		fmt.Println("es 1")

	} else {
		fmt.Println("no es 1 ")
	}
	//with and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("es correcto")
	} else {
		fmt.Println("nada no es ni 1 ni 2")
	}

	//with or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("es verdad")

	}
	//password
	if pass == 123 && email == "correo" {
		fmt.Println("iniciando sesión")
	} else {
		fmt.Println("error login")
	}

}
