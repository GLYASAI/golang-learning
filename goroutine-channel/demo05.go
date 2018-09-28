package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("%d ", <-ch)
	fmt.Printf("%d ", <-ch)
	fmt.Printf("%d ", <-ch)
}
