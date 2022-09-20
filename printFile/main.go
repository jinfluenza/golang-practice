package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) *os.File {
	o, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Here is the error message: %s", err)
		os.Exit(1)
	}
	return o
}

func main() {
	o := readFile(os.Args[1])
	// fmt.Println(o)

	io.Copy(os.Stdout, o)
}
