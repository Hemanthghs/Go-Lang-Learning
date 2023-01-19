package main

import (
	"context"
	"fmt"
	"time"
)

func sleepAndTalk2(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	sleepAndTalk2(ctx, 5*time.Second, "Hemanthsai")
}
