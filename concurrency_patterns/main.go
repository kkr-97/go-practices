package main

import (
	"fmt"
	"sync"
)

// ######concurrency - 1

// func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
// 	stream := make(chan T)
// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case stream <- fn():
// 			}
// 		}
// 	}()
// 	return stream
// }

// func main() {
// 	done := make(chan int)
// 	defer close(done)

// 	fn := func() int { return rand.Intn(1000) }

// 	for val := range repeatFunc(done, fn) {
// 		fmt.Println(val)
// 	}
// }

//CONCURRENCY - 2

var wg sync.WaitGroup

func main() {

	done := make(chan any)
	defer close(done)

	// done <- "kkr"
	// done <- 7

	// fmt.Println(<-done)
	// fmt.Println(<-done)

	cows := make(chan any, 30)
	pigs := make(chan any, 30)

	go func() {
		for {
			select {
			case <-done:
				return
			case cows <- "moo":

			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case pigs <- "mink":

			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows)
	wg.Add(1)
	go consumePigs(done, pigs)

	wg.Wait()

}

func orDone(done <-chan any, absChan <-chan any) <-chan any {
	inpStream := make(chan any)
	go func() {
		defer close(inpStream)
		for {
			select {
			case <-done:
				return
			case val, ok := <-absChan:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				// inpStream <- val
				select {
				case inpStream <- val:
				case <-done:
					return
				}

			}
		}
	}()
	return inpStream
}

func consumeCows(done <-chan any, cows <-chan any) {
	defer wg.Done()
	for cow := range orDone(done, cows) {
		fmt.Println(cow)
	}
}

func consumePigs(done <-chan any, pigs <-chan any) {
	defer wg.Done()
	for pig := range orDone(done, pigs) {
		fmt.Println(pig)
	}
}
