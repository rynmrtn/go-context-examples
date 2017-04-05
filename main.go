package main

import (
	"context"
	"fmt"
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

	ctx, _ := context.WithTimeout(context.Background(), 10000)
	go sayWithContext(ctx, "context")

	select {
	case <-ctx.Done():
		fmt.Println("done")
	}
}
