package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

// Read the folder
func readFolder() []string {
	path := "notes/"
	var outputName string
	var outputArr []string
	files, err := os.ReadDir(path)
	check(err)
	for _, file := range files {
		outputName = strings.TrimSuffix(file.Name(), ".txt")
		outputArr = append(outputArr, outputName)
	}
	outputArr = append(outputArr, "Create new file")
	outputArr = append(outputArr, "Exit")
	return outputArr
}

// same as a menu but choose the file to work with it
func chooseFile(options []string, selected int) {
	for i, option := range options {
		if selected == len(options)-2 && i == len(options)-2 {
			fmt.Printf("\033[30;45m%s\033[0m\n", option)
		} else if i == len(options)-2 {
			fmt.Printf("\033[35m%s\033[0m\n", option)
		} else if selected == len(options)-1 && i == len(options)-1 {
			fmt.Printf("\033[30;41m%s\033[0m\n", option)
		} else if i == len(options)-1 {
			fmt.Printf("\033[31m%s\033[0m\n", option)
		} else if i == selected {
			fmt.Printf("\033[30;46m%s\033[0m\n", option)
		} else {
			fmt.Printf("\033[36m%s\033[0m\n", option)
		}
	}
}

// keyboard arrows to navigate
func fileSelect() string {
	options := readFolder()
	selected := 0

	err := keyboard.Open()
	check(err)
	if len(options) < 1 {
		fmt.Println("\033[33mYou don't have any notes\033[0m")
		return ""
	}
	chooseFile(options, selected)
	for {

		_, key, err := keyboard.GetKey()
		check(err)
		switch key {

		case keyboard.KeyArrowUp:
			selected--
			if selected < 0 {
				selected = len(options) - 1
			}

		case keyboard.KeyArrowDown:
			selected++
			if selected >= len(options) {
				selected = 0
			}

		case keyboard.KeyEnter:
			if selected == len(options)-2 {
				keyboard.Close()
				name := CreateFile()
				return name
			} else if selected == len(options)-1 {
				EmptyTerminal()
				fmt.Printf("\n\033[42m	===Goodbye!===		\033[0m\n")
				os.Exit(0)
			} else {
				keyboard.Close()
				return options[selected]
			}

		default:
			continue
		}
		cursorOff()
		fmt.Printf("\033[%dA", len(options))
		chooseFile(options, selected)
	}
}
