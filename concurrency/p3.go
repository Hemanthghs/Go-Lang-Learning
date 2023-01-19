package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func sleepAndTalk(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}

// func main() {
// 	ctx := context.Background()
// 	sleepAndTalk(ctx, 5*time.Second, "Hemanth")
// }

// func main() {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, time.Second)
// 	defer cancel()
// 	sleepAndTalk(ctx, 5*time.Second, "Hemanth")
// }

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	sleepAndTalk(ctx, 5*time.Second, "Hemanth")
}
