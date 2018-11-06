package context

import (
	"context"
	"fmt"
	"time"
)

func cancelTest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("Done...\n")
			return
		default:
			fmt.Printf("Working...\n")
		}
		time.Sleep(1 * time.Second)
	}
}

func closeChan(stop chan struct{}) {
	select {
	case <-stop:
		fmt.Print("execute...")
	}
}
