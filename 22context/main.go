package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	defer cancel()

	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			println("This routine cancelled:", ctx.Err())
		}

	}(ctx, &wg)

	fmt.Println("main func exited...")

	wg.Wait()
}
