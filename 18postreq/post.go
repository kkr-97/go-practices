package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myUrl = "http://localhost:3001/user_login/"

func jsonPost() {
	//json payload
	reqBody := strings.NewReader(`
		{
			"username":"JoeBiden",
			"password":"biden@123"
		}
	`)

	resp, err := http.Post(myUrl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("status code:", resp.StatusCode)
	fmt.Println("response:", string(respBody))
}

func main() {
	//performing POST request
	// jsonPost()

}
