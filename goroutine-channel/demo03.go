package main

import "fmt"

var complete = make(chan int)

func loop03() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	complete <- 0
}

func main() {
	go loop03()

	<- complete
}
