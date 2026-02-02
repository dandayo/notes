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
	var fileName string
	if os.Args[1] == "--help"{
		Help()
	} else {
		fileName = os.Args[1]
		OpenFile(fileName)
		Menu(fileName)
	}
}