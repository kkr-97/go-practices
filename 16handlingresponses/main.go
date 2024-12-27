package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func getCode(resCodeCh chan int) {
	resp, _ := http.Get(url)
	fmt.Printf("type of response object: %T\n", resp) // Prints *http.Response
	defer resp.Body.Close()
	resCodeCh <- resp.StatusCode
}

func case2() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bodyContent, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Type of response body: %T\n", bodyContent) //prints []uint8 --> bytecode
	fmt.Println("response body content:", string(bodyContent))

}

func main() {
	// resCodeCh := make(chan int)
	// go getCode(resCodeCh)
	// fmt.Println("before printing status code")
	// fmt.Println(<-resCodeCh) //blocks the go routine
	// fmt.Println("after printing status code")

	// case 2: reading body content from response
	case2()
}
