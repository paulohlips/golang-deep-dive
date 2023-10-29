package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("newfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := f.Write([]byte("Creating new file \n"))

	fmt.Printf("File created! length: %d bytes", length)

	f.Close()

	file, err := os.ReadFile("newfile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(file))
}
