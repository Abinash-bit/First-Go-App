package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)

	}
	fmt.Printf("Hello World\nos.Args: %v\nArgument: %v\n", args, args[1])

}
