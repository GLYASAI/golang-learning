package main

import "fmt"

func main() {
	arr1 := [...]string {"alpha", "beta", "gamma"}
	arr2 := [3]string{}

	arr2 = arr1

	arr2[0] = "alpha2"
	fmt.Println(arr1[0])
	fmt.Println(arr2[0])
}

