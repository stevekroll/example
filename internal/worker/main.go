package main

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

type ctxKey string

const (
	levelKey ctxKey = "level"
)

func printVersion() {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	print(string(out))
}

func run() {
	// Create a root context with a timeout.
	ctx := context.WithValue(context.Background(), levelKey, "main")
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	// Schedule a function to run after the context is done.
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("cancel origin...", ctx.Value(levelKey))
	})
	defer stop()

	// Start a worker goroutine.
	var wg sync.WaitGroup
	wg.Go(func() {
		name := "worker A"
		worker(ctx, name, nil)
	})
	wg.Go(func() {
		name := "worker B"
		timeout := 3 * time.Second
		worker(ctx, name, &timeout)
	})
	wg.Go(func() {
		name := "worker C"
		timeout := 7 * time.Second
		worker(ctx, name, &timeout)
	})
	wg.Go(func() {
		name := "worker D"
		workerWithoutTimeout(ctx, name)
	})

	// Wait for the worker to finish.
	wg.Wait()
}

func worker(ctx context.Context, name string, timeout *time.Duration) {
	ctx = context.WithValue(ctx, levelKey, name)

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	t := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			// Clean up resources if necessary.
			fmt.Println("context cancelled...", ctx.Value(levelKey))
			return
		case <-t.C:
			// Do some work.
			fmt.Println("tick...", ctx.Value(levelKey))
		}
	}
}

func workerWithoutTimeout(ctx context.Context, name string) {
	ctx = context.WithoutCancel(context.WithValue(ctx, levelKey, name))

	t := time.NewTicker(time.Second)
	timer := time.NewTimer(time.Second * 15)
	for {
		select {
		case <-timer.C:
			fmt.Println("worker without timeout finished...", ctx.Value(levelKey))
			return
		case <-ctx.Done():
			// Clean up resources if necessary.
			fmt.Println("context cancelled...", ctx.Value(levelKey))
			return
		case <-t.C:
			// Do some work.
			fmt.Println("tick...", ctx.Value(levelKey))
		}
	}
}

func main() {
	printVersion()
	run()
}
