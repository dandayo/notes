package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function for user input
func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Printf("\033[38;5;60m->\033[0m")
	scanner.Scan()
	answer := strings.TrimSpace(scanner.Text()) //Clean the input from \n \t and space
	if answer == "" {
		fmt.Println("\033[91mEmpty input.\nTry again!\033[0m") //Check empty input
	}
	if answer == "--help" {
		Help() //Call help
	}
	return answer
}

// Clean the terminal
func EmptyTerminal() {
	fmt.Print("\033[3J\033[H\033[2J")
}

// Check the errors
func check(e error) {
	if e != nil {
		fmt.Println("\033[91mMistake! Something gone wrong\033[0m")
	}
}
