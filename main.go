package main

import (
    "fmt"
    "os"

)

func main() {
	EmptyTerminal()
	fmt.Println("\033[32m\n===Welcome to the notes tool!===\033[0m")
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./todotool [TAG]\nAlso use --help to read about this tool")
		return
	}

	if os.Args[1] == "--help"{
		Help()
	} else {
		OpenFile(os.Args[1])
		Menu(os.Args[1])
	}
}
