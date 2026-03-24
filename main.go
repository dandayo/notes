package main

import (
	"fmt"
	"os"
)

// Call the main function that start the tool
func main() {
	EmptyTerminal() //Clean terminal
	fmt.Println("\n\033[32m===Welcome to the notes tool!===\033[0m\n")
	password := CreatePassword() //Create or check password
	key := []byte(password)

	err := os.MkdirAll("notes", 0700)
	check(err)

	fmt.Println("\033[33mChoose your file or create a new:\n\033[0m")
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		Help()
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "--reset" {
		fmt.Println("\033[33mWrite your password to confirm to reset it\n\033[0m")
		updatePassword()
		return
	}
	cursorOff()
	filename := fileSelect()
	EmptyTerminal()     // to clean after open file
	Menu(filename, key) //Start menu
}
