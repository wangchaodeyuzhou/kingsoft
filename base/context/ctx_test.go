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

func TestExampleWithCancel(t *testing.T) {

	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}

		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestExampleWithDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestExampleWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestExampleWithValue(t *testing.T) {
	type favContextKey string

	// 一个查找而已
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}

		fmt.Println("key not found: ", k)
	}

	k := favContextKey("wrc")
	ctx := context.WithValue(context.Background(), k, "wrc nb")
	f(ctx, k)
	f(ctx, favContextKey("color"))
}

func TestContextSendSubThread(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "任务1")
	go watch(ctx, "任务2")
	go watch(ctx, "任务3")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知任务也停止")
	defer cancel()
	time.Sleep(5 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "任务退出，停止")
			return
		default:
			fmt.Println(name, "goroutine 任务中...")
			time.Sleep(2 * time.Second)
		}
	}
}
