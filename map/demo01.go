package main

import "fmt"

func main() {
	girlMap := map[string]int{"yp": 0, "sy": 1}

	for key, value := range girlMap {
		fmt.Printf("name: %s, %d\n", key, value)
	}
}
