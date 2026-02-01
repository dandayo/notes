package main

import (
    "fmt"
    "os"

    "github.com/manifoldco/promptui"
)

func Menu(name string) {
	msg := ""
	items := []string{
		"Show notes",
		"Add a note",
		"Delete a note",
		"Exit",
	}

	for msg != "4" {

		prompt := promptui.Select{
			Label: "Select operation",
			Items: items,
			Size: 4,
		}

		index, _, err := prompt.Run()
		check(err)

		msg = input()

        switch msg {

            case "1":
            	EmptyTerminal()
            	fmt.Println("\033[1m===Your notes!===\n\033[0m")
                ReadFile(name)
            case "2":
            	fmt.Print("\033[3J\033[H\033[2J")
            	fmt.Println("\033[34mEnter the note text:\033[0m")
                note := input()
			    AddNote(name, note)

            case "3":
           		EmptyTerminal()
            	fmt.Println("\033[33mEnter the number of note to remove or 0 to cancel:\033[0m")
                index := input()
                if index == "0" {
                continue
                } else {
                	RemoveNote(name, index)
                }

            case "4":
           		EmptyTerminal()
            	fmt.Println("\033[1m===Goodbye!===\n\033[0m")
			    os.Exit(0)

            default:
            	if msg != "" && msg !="--help"{
	             	fmt.Println("\033[31mYou have entered an invalid input!\033[0m")
				    continue
             	}


        }
	}
}
