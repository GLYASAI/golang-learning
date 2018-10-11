package main

import "fmt"

func main() {
	channels := []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1), // 最后的","不能省略
	}

	//channels[0] <- 1
	//channels[1] <- 1
	//channels[2] <- 1

	select {
	case <-channels[0]:
		fmt.Println("0")
	case <-channels[1]:
		fmt.Println("1")
	case <-channels[2]:
		fmt.Println("2")
	case channels[0] <- 2:
		fmt.Println(4)
	default:
		fmt.Println("default")
	}
}
