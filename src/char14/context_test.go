package char14

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCacelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelled(t *testing.T) {
	ctx, cacelled := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCacelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "is Done")
		}(i, ctx)
	}
	cacelled()
	time.Sleep(time.Second*1)
}
