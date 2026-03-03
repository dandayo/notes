package main

import (
	"fmt"
	"os"
)

// Call the main function that start the tool
func main() {
	EmptyTerminal() //Clean terminal
	fmt.Println("\n\033[32m===Welcome to the notes tool!===\033[0m\n")
	fmt.Println("\033[33mChoose your file or create a new:\n\033[0m")
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		Help()
		return
	}
	cursorOff()
	filename := fileSelect()
	EmptyTerminal() // to clean after open file
	fmt.Printf("\033[4mTitle: %s\033[0m\n", filename)
	Menu(filename) //Start menu
}
