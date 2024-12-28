package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myUrl = "https://google.com"

func getCode(resCodeCh chan int) {
	resp, _ := http.Get(myUrl)
	fmt.Printf("type of response object: %T\n", resp) // Prints *http.Response
	defer resp.Body.Close()
	resCodeCh <- resp.StatusCode
}

func case2() {
	resp, err := http.Get(myUrl)
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
	// case2()

	// case 3: Handling myUrls
	const myUrl = "https://google.com:3030/keerthan-para?auth=false"
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //google.com:3030
	fmt.Println(result.Path)     ///keerthan-para
	fmt.Println(result.Port())   //3030
	fmt.Println(result.RawQuery) //auth=false
	//queryparameters
	qParams := result.Query() //type:url.Values, value: map[auth:[false]]
	fmt.Printf("type:%T, value: %v", qParams, qParams)
}
