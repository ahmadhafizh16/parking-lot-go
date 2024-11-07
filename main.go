package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file with commands.")
		return
	}

	filename := os.Args[1]

	commandProcessor := &CommandProcessor{}
	commandProcessor.ProcessCommands(filename)
}
