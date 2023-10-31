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

	res, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	var newCustomer Customer
	json.Unmarshal(res, &newCustomer)
	println(newCustomer.Name)
}
