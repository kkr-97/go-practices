package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Adding new content

	newContent := "Keerthan, this is a text that is not written, but injected."
	file, err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, newContent)
	if err != nil {
		panic(err)
	}

	fmt.Println("file created successfully with length:", length)

	// trying to read content from file
	// for that we use ioutil module, and the data returned will be in the form of bytes
	byteData, err := os.ReadFile("newFile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Content in the file:\n", string(byteData))

	defer file.Close()
}
