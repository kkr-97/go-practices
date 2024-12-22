package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["js"] = "javascrip"
	languages["py"] = "python"
	languages["cpp"] = "c++"

	fmt.Println("List of languages: ", languages)

	// Iterating over map
	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// deleting "py"
	delete(languages, "py")
	fmt.Println("List of languages after deleting 'py': ", languages)

	// clearing values
	clear(languages)
	fmt.Println("List of languages after clearing values: ", languages)
}
