package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func sleepTalk(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Fatal(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	sleepTalk(ctx, 1*time.Second, "Hemanth")

}
