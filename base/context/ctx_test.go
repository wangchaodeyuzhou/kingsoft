package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestHandlerRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go HandlerRequest(ctx)

	time.Sleep(5 * time.Second)

	fmt.Println("it is time to stop all sub goroutine")
	cancel()

	time.Sleep(5 * time.Second)
}

func TestHandlerRequestTimeOut(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandlerRequest(ctx)

	time.Sleep(10 * time.Second)
}

func TestHandlerRequestValueCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), "p", 1)
	go HandlerRequest(ctx)

	time.Sleep(10 * time.Second)
}
