package context

import (
	"context"
	"testing"
	"time"
)

func TestCancelTest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go cancelTest(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func TestCloseChan(t *testing.T) {
	stop := make(chan struct{})
	go closeChan(stop)

	time.Sleep(1 * time.Second)

	//go func(stop chan struct{}) {
	//	stop<-struct{}{}
	//	fmt.Print("send")
	//}(stop)

	close(stop)


	time.Sleep(3 * time.Second)
}
