package main

import "fmt"

func normalTayson(message string) {
	fmt.Println(message)
	fmt.Println("te amo mucho")
	fmt.Println("Dios es bueno")

}

func taysonTres(a int, b int, c string) {
	fmt.Println(a, b, c)

}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalTayson("hola v")
	taysonTres(1, 2, "holi")

	value := returnValue(2)
	fmt.Println("value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("valor 1 y 2:", value1, "-", value2)

}
