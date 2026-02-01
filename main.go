package main

import (
    "fmt"
    "os"

)

func main() {
	EmptyTerminal()
	fmt.Println("\033[32m\n===Welcome to the notes tool!===\033[0m")
	fileName := os.Args[1]
	if fileName == "" {
		 fmt.Println("Usage: ./todotool [TAG]\nAlso use --help to read about this tool")
	}
	OpenFile(fileName)
	Menu(fileName)
}
