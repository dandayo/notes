package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func readKey() keyboard.Key {
	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	return key
}

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
			fmt.Printf("%s%d. %s\033[0m\n", oppositeColors[i], i+1, option)
		} else {
			fmt.Printf("%s%d. %s\033[0m\n", colors[i], i+1, option)
		}
	}

	fmt.Println()
}

func Menu(name string) {
	selected := 0

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		renderMenu(selected)
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {

		case keyboard.KeyArrowUp:
			selected--
			EmptyTerminal()
			if selected < 0 {
				selected = 4
			}

		case keyboard.KeyArrowDown:
			selected++
			EmptyTerminal()
			if selected > 4 {
				selected = 0
			}

		case keyboard.KeyEnter:
			keyboard.Close()
			execute(selected, name)
			keyboard.Open()
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
		fmt.Println("\033[33mEnter the number of note to remove or 0 to cancel:\033[0m")
		index := input()
		if index == "0" {
			return
		} else if index == "--help" {
			return
		} else {
			RemoveNote(name, index)
		}
		fmt.Println("\n")
		ReadFile(name)

	case 3:
		EmptyTerminal()
		fmt.Println("\033[91mAre you sure?! Press '1' if agree to delete file\033[0m")
		doublecheck := input()
		if doublecheck != "1" {
			return
		} else {
			DeleteFile(name)
			os.Exit(0)
		}

	case 4:
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
}
