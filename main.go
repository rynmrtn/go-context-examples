package main

import (
	"context"
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
	}
}

func sayWithContext(ctx context.Context, s string) {
	for i := 0; i < 1000000; i++ {
		fmt.Println(s, i)
	}
}

func main() {
	fmt.Println("Starting go routine")

	say("good morning")

	go say("hey")
	d, _ := time.ParseDuration("3s")
	ctx, _ := context.WithTimeout(context.Background(), d)
	go sayWithContext(ctx, "me")
	go sayWithContext(ctx, "first")

	select {
	case <-ctx.Done():
		fmt.Println("done")
	}
}
