package main

import "fmt"

func main() {
	ch := make(chan int)

	for c := range ch {
		fmt.Println(c)
	}
	fmt.Println("end")
}
