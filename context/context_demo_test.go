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
