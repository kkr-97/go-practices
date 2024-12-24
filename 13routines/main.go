package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

func printStatusCode(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer resp.Body.Close()

	statusCode := resp.StatusCode
	fmt.Println(url, " --> ", statusCode) // prints the status code of the HTTP response
}

func main() {

	// go greeter("first")
	// greeter("second")

	// go num := 1 //not possible, because go is only for functions
	// fmt.Println(num)

	webLists := []string{
		"https://www.google.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"SIAVIA",
	}

	// for _, url := range webLists {
	// 	//below takes some time to execute
	// 	printStatusCode(url)
	// }

	//below will execute but not wait for the routine to return the control before main thread deads.
	// for _, url := range webLists {
	// 	go printStatusCode(url)
	// }

	//using Wait group -> Add(), Done(), Wait()

	for _, url := range webLists {
		wg.Add(1)
		go printStatusCode(url)

	}
	wg.Wait() // this will not allow main to complete till wait grp does gives signal

}
