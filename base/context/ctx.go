package context

import (
	"context"
	"fmt"
	"time"
)

func HandlerRequest(ctx context.Context) {
	go WithRedis(ctx)
	go WithDataBase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handler request done")
			return
		default:
			fmt.Println("handler request running", ctx.Value("p"))
			time.Sleep(2 * time.Second)
		}
	}

}

func WithRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("with redis done")
			return
		default:
			fmt.Println("With redis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WithDataBase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("with database done")
			return
		default:
			fmt.Println("with database running")
			time.Sleep(2 * time.Second)
		}
	}

}
