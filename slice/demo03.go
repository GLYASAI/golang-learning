package main

import "fmt"

func main() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	slice := source[2:3:3]

	slice = append(slice, "Kiwi")

	for _, v := range source {
		fmt.Printf("%s\n", v)
	}
}
