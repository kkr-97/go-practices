package main

import (
	"fmt"
	"sync"
)

func example1() {
	ch := make(chan int)
	var ans int

	func() {
		sum := 0
		for i := 1; i < 10; i++ {
			sum += i
		}
		ch <- sum
		ans = sum

	}()

	fmt.Println(ans)
	op := <-ch
	fmt.Println(op, ans)
}

func example2() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		fmt.Println(<-ch) // block the code, till the value in the channel is received
		fmt.Println("channel value received")
		wg.Done()
	}()

	go func() {
		ch <- 5
		wg.Done()
	}()

	wg.Wait()
}

func example3() {
	ch := make(chan int, 2)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		fmt.Println(<-ch) // blocks the code, till the value in the channel is received
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)

		fmt.Println("channel value received")
		wg.Done()
	}()

	go func() {
		ch <- 5
		ch <- 666 //trying to reassign the channel value, which will throws an error
		ch <- 444 //If a sender tries to send a value when the buffer is full, it will block until a receiver removes a value from the buffer.
		ch <- 555
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	// example1()
	// example2()
	example3()

	// ch2 := make(chan string) //unbuffered channel

	// //below gives runtime error because, An unbuffered channel requires both a sender and a receiver to be ready at the same time.
	// ch2 <- "hello"
	// fmt.Println(<-ch2)
}
