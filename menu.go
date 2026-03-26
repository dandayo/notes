package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func renderMenu(selected int) {
	options := []string{
		"Show notes",
		"Add a note",
		"Delete a note",
		"Delete all notes. Be careful!",
		"Help",
		"Exit",
	}

	colors := []string{
		"\033[36m",
		"\033[34m",
		"\033[33m",
		"\033[31m",
		"\033[32m",
		"\033[35m",
	}

	oppositeColors := []string{
		"\033[30;46m",
		"\033[30;44m",
		"\033[30;43m",
		"\033[30;41m",
		"\033[30;42m",
		"\033[30;45m",
	}

	fmt.Println("\nSelect operation:\n")

	for i, option := range options {
		if i == selected {
			fmt.Printf("->%s %s\033[0m\n", oppositeColors[i], option)
		} else {
			fmt.Printf("%s%d. %s\033[0m\n", colors[i], i+1, option)
		}
	}

	fmt.Println()
}

func Menu(name string, password []byte) {
	selected := 0
	var height int = 10 // lines in menu

	err := keyboard.Open()
	check(err)
	defer keyboard.Close()

	renderMenu(selected)
	for {
		_, key, err := keyboard.GetKey()
		check(err)
		cursorOff()
		switch key {

		case keyboard.KeyArrowUp:
			selected--
			if selected < 0 {
				selected = 5
			}
			clearMenu(height)
			renderMenu(selected)

		case keyboard.KeyArrowDown:

			selected++
			if selected > 5 {
				selected = 0
			}
			clearMenu(height)
			renderMenu(selected)

		case keyboard.KeyEnter:
			keyboard.Close()
			execute(selected, name, password)
			keyboard.Open()
			renderMenu(selected)

		default:
			clearMenu(height)
			renderMenu(selected)
		}
	}
}

func execute(selected int, name string, password []byte) {
	switch selected {

	case 0:
		EmptyTerminal()
		fmt.Printf("\033[4mTitle: %s\033[0m\n", name)
		ReadFile(name, password)

	case 1:
		EmptyTerminal()
		fmt.Println("Enter the note text:")
		cursorOn()
		note := Input()
		AddNote(name, note, password)
		cursorOff()

	case 2:
		EmptyTerminal()
		ReadFile(name, password)
		fmt.Printf("\n")
		fmt.Println("\033[33mEnter the number of note to remove or 0 to cancel:\033[0m")
		cursorOn()
		index := Input()
		if index == "0" {
			return
		} else if index == "--help" {
			return
		} else {
			RemoveNote(name, index, password)
		}
		cursorOff()
		EmptyTerminal()

	case 3:
		EmptyTerminal()
		fmt.Println("\033[91mAre you sure?! Press '1' if agree to delete file\033[0m")
		cursorOn()
		doubleCheck := Input()
		if doubleCheck != "1" {
			return
		} else {
			DeleteFile(name)
			os.Exit(0)
		}
		cursorOff()

	case 4:
		EmptyTerminal()
		Help()

	case 5:
		EmptyTerminal()
		fmt.Printf("\n\033[42m	===Goodbye!===		\033[0m\n")
		os.Exit(0)
	}
}
