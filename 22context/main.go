package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func withCancelExample() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	defer cancel()

	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		select {
		case <-ctx.Done():
			println("This routine cancelled:", ctx.Err())
			wg.Done()
		}

	}(ctx, &wg)

	fmt.Println("main func exited...")

	wg.Wait()
}

func routineGo(ctx context.Context, wg *sync.WaitGroup, res chan<- int) {
	defer wg.Done()
	fmt.Println("go routine started")
	select {

	case <-time.After(3 * time.Second):
		fmt.Println("If this prints, ctx is still running")
	case <-ctx.Done():
		fmt.Println("Routine timeout")
		res <- -1
		return
	}

	fmt.Println("rest of go routing code")
	res <- 1
}

func withTimeOutExample() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	var wg sync.WaitGroup
	res := make(chan int)

	wg.Add(1)
	go routineGo(ctx, &wg, res)

	fmt.Println("Result Status:", <-res)

	fmt.Println("End of main func")
	wg.Wait()

}

func main() {
	// withCancelExample()

	withTimeOutExample()

}
