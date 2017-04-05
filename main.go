package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sayWithContext(ctx context.Context, s string) {
	for i := 0; i < 1000000; i++ {
		fmt.Println(s, i)
	}
}

func main() {
	signalChannel := make(chan struct{})
	d, _ := time.ParseDuration("2s")
	pctx, cancel := context.WithTimeout(context.Background(), d)

	type kv string

	f := func(ctx context.Context, k kv) {
		sleepDur, _ := time.ParseDuration("1s")
		if v := ctx.Value(k); v != nil {
			fmt.Println("val: ", v)
			time.Sleep(sleepDur)
			return
		}
		fmt.Println("key not found")
	}

	// Listen for INT and SIGTERM syscalls, clean up
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		<-sig
		// Close the context
		cancel()
		close(signalChannel)
	}()

	go sayWithContext(pctx, "me")
	go sayWithContext(pctx, "first")

	k := kv("lang")
	vctx := context.WithValue(pctx, k, "go")

	go f(vctx, k)

	select {
	case <-pctx.Done():
		if pctx.Err() == context.DeadlineExceeded {
			fmt.Println("Expired")
		} else if pctx.Err() == context.Canceled {
			fmt.Println("Canceled")
		}
	}
}
