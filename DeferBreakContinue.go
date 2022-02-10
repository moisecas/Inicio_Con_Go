package main

import "fmt"

func main() {
	//defer
	defer fmt.Println("Hola")
	fmt.Println("mundo")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("es 2")
			continue
		}
		//break
		if i == 7 {
			fmt.Println("break")
			break
		}

	}

}
