package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["moises"] = 30

	fmt.Println(m)

	//recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//encontrar un valor
	value, ok := m["jose"]
	fmt.Println(value, ok)
}
