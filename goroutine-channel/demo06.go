package main

import "fmt"

func init() {
	fmt.Println("call init()")
}

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // close之后，只能读，不能写

	for v := range (ch) {
		fmt.Printf("%d ", v)
	}
}
