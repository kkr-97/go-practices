package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//working with twitter backend server in localhost:3001

	const myUrl = "http://localhost:3001/"

	resp, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("status code: ", resp.StatusCode)
	fmt.Println("Protocol: ", resp.Proto)

	content, _ := io.ReadAll(resp.Body)
	fmt.Println("content:", string(content))

}
