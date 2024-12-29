package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name     string      `json:"name"`          //aliases
	Password string      `json:"-"`             //doesn't show the password field whoever consume this json data
	Age      interface{} `json:"age,omitempty"` //if value is nil, don't throw this field
}

func EncodeJson() {
	students := []student{
		{"John", "123", 20},
		{"Alice", "456", 22},
		{"Bob", "789", nil},
	}

	//packing the json
	jsonData, err := json.MarshalIndent(students, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of jsondata: %T\n", jsonData) //[]uint8
	fmt.Println("json data: \n", string(jsonData))

	DecodeJson(jsonData)
}

func DecodeJson(jsonData []byte) {
	// case1: using structs
	// var students []student
	//validating JSON
	isValidJson := json.Valid(jsonData)
	if !isValidJson {
		panic("invalid json")
	}
	// //unmarshalling the JSON
	// json.Unmarshal(jsonData, &students)
	// fmt.Printf("after unmarshalling: %v\n", students)      // [{John  20} {Alice  22} {Bob  <nil>}]
	// fmt.Printf("Type after unmarshalling: %T\n", students) // []main.student

	//case2: key, value extraction
	var studentsMap []map[string]interface{}
	json.Unmarshal(jsonData, &studentsMap)
	fmt.Printf("After unmarshalling to map: %v\n", studentsMap)
	//studentsMap is a map of string to interface{} (key, value)
	for _, val := range studentsMap {
		for key, value := range val {
			fmt.Printf("key: %v, value: %v\n", key, value)
		}
		fmt.Println("")
	}

}

func main() {
	// creating json data
	EncodeJson() //decoding is also done by composition
}
