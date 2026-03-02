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
		"Exit",
	}

	colors := []string{
		"\033[36m",
		"\033[34m",
		"\033[33m",
		"\033[31m",
		"\033[35m",
	}
	oppositeColors := []string{
		"\033[30;46m",
		"\033[97;44m",
		"\033[30;43m",
		"\033[97;41m",
		"\033[97;45m",
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

func clearMenu(height int) {
	fmt.Printf("\033[%dA", height) //
	for i := 0; i < height; i++ {
		fmt.Print("\033[2K\n")
	}
	fmt.Printf("\033[%dA", height)
}

func Menu(name string) {
	selected := 0
	var height int = 9 // lines in menu
	err := keyboard.Open()
	check(err)

	defer keyboard.Close()
	renderMenu(selected)
	for {
		_, key, err := keyboard.GetKey()
		check(err)

		switch key {

		case keyboard.KeyArrowUp:
			selected--
			if selected < 0 {
				selected = 4
			}
			clearMenu(height)
			renderMenu(selected)

		case keyboard.KeyArrowDown:

			selected++
			if selected > 4 {
				selected = 0
			}
			clearMenu(height)
			renderMenu(selected)

		case keyboard.KeyEnter:
			keyboard.Close()
			execute(selected, name)
			keyboard.Open()
			renderMenu(selected)

		default:
			clearMenu(height)
			renderMenu(selected)
		}
	}
}

func execute(selected int, name string) {
	switch selected {

	case 0:
		EmptyTerminal()
		fmt.Printf("\033[4mTitle: %s\033[0m\n", name)
		ReadFile(name)

	case 1:
		EmptyTerminal()
		fmt.Println("Enter the note text:")
		note := input()
		AddNote(name, note)

	case 2:
		EmptyTerminal()
		ReadFile(name)
		fmt.Println("\n")
		fmt.Println("\033[33mEnter the number of note to remove or 0 to cancel:\033[0m")
		index := input()
		if index == "0" {
			return
		} else if index == "--help" {
			return
		} else {
			RemoveNote(name, index)
		}
		EmptyTerminal()

	case 3:
		EmptyTerminal()
		fmt.Println("\033[91mAre you sure?! Press '1' if agree to delete file\033[0m")
		doubleCheck := input()
		if doubleCheck != "1" {
			return
		} else {
			DeleteFile(name)
			os.Exit(0)
		}

	case 4:
		EmptyTerminal()
		fmt.Println("\n\033[42mGoodbye!\033[0m\n")
		os.Exit(0)
	}
}
