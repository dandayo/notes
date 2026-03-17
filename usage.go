package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

// Function for user input
func Input() string {
	keyboard.Close()
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

// Clean the menu
func clearMenu(height int) {
	fmt.Printf("\033[%dA", height) //
	for i := 0; i < height; i++ {
		fmt.Print("\033[2K\n")
	}
	fmt.Printf("\033[%dA", height)
}

// Check the errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cursorOn() {
	fmt.Print("\033[?25h")
}

func cursorOff() {
	fmt.Print("\033[?25l")
}
