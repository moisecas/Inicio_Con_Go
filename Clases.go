package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "ford", year: 2020}
	fmt.Println(myCar)

	//OTRA FORMA
	var otherCar car
	otherCar.brand = "ferrari"
	otherCar.year = 2021
	fmt.Println(otherCar)

}
