package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice[0] = 100
	fmt.Println(slice[1])
}
