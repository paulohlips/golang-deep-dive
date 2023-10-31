package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Id   int
	Name string
}

func main() {
	customer := Customer{Id: 1, Name: "James Bond"}

	json, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(json))
}
