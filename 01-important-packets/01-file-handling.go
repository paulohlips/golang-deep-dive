package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "new-file.txt"
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	length, err := f.Write([]byte("Creating new file big file to print using buffer\n"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created! length: %d bytes", length)

	f.Close()

	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(file))

	// using buffers"newFile.txt"
	bigFile, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(bigFile)
	buffer := make([]byte, 1)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}
}
