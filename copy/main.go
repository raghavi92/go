package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args[1])

	file, err := os.Open(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, file)
	}
}
