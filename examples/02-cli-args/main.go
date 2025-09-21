package main

import (
	"fmt"
	"os"
)

func main() {
	// Option 1 to declare: var args []string
	// Option 2 to declare: args := variable
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: ./hello-world <argument>\n")
		os.Exit(1)
	}
	fmt.Printf("Hello, World!\nos.Args: %v\nArgument: %v\n", args, args[1:])
}
