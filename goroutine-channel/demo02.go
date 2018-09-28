package main

import "fmt"

func main() {
	messages := make(chan string)
	go func(message string) {
		messages <- message
	}("Ping!")

	fmt.Printf(<-messages)
}
