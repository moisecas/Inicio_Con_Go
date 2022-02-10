package main

import "fmt"

func say(text int, c chan<- int) {
	c <- text
}

func main() {
	c := make(chan int, 1)
	fmt.Println("hola")
	go say(2, c)
	fmt.Println(<-c)
}
