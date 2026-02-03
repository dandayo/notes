package main

import (
    "fmt"
    "os"

)

//Call the main function that start the tool
func main() {
	EmptyTerminal() //Clean terminal
	fmt.Println("\033[32m\n===Welcome to the notes tool!===\033[0m")
	if len(os.Args) != 2 { //Check is it empty input
		fmt.Println("Usage: ./todotool [TAG]\nAlso use --help to read about this tool")
		return
	}

	if os.Args[1] == "--help"{
		Help()
	} else {
		OpenFile(os.Args[1])
		Menu(os.Args[1])//Start menu
	}
}
