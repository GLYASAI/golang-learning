package main

import "fmt"

func main() {
	ch := getIntChan()
	for c := range ch {
		fmt.Println(c)
	}
}

func getIntChan() <-chan int {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
