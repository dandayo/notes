package main

import (
    "fmt"
    "os"
)
//Menu func
func Menu(name string) {
	msg := ""
	for msg != "5"{
		//Print for choose
		fmt.Println("\nSelect operation:\n")
		fmt.Println("\033[36m1. Show notes.\033[0m")
		fmt.Println("\033[34m2. Add a note.\033[0m")
		fmt.Println("\033[33m3. Delete a note.\033[0m")
		fmt.Println("\033[31m4. Delete all notes. Be careful\033[0m")
		fmt.Println("\033[35m5. Exit.\033[0m")
		fmt.Println("\n")
		msg = input()

        switch msg {

            case "1": //Show the list of note
            	EmptyTerminal()
                fmt.Printf("\033[4mTitle: %s\033[0m\n", name)
                ReadFile(name)
            case "2"://Add new note
            	fmt.Print("\033[3J\033[H\033[2J")
            	fmt.Println("\033[34mEnter the note text:\033[0m")
                note := input()
                if note == "--help"{
                	continue
                } else {
                	AddNote(name, note)
                }

            case "3"://Delete the note
           		EmptyTerminal()
            	fmt.Println("\033[33mEnter the number of note to remove or 0 to cancel:\033[0m")
                index := input()
                if index == "0" {
                continue
                } else if index == "--help"{
                	continue
                } else {
                	RemoveNote(name, index)
                }
            case "4" ://Delete the file with notes
            	EmptyTerminal()
                fmt.Println("\033[91mAre you sure?! Press '1' if agree to delete file\033[0m")
             	doublecheck := input()
             	if doublecheck != "1"{
              		continue
              	} else {
                    DeleteFile(name)
                    os.Exit(0)
               }

            case "5"://End process
           		EmptyTerminal()
            	fmt.Println("\033[1m===Goodbye!===\n\033[0m")
			    os.Exit(0)
				EmptyTerminal()

            default://Invalid input
            	if msg != "" && msg !="--help"{
	             	fmt.Println("\033[91mYou have entered an invalid input!\033[0m")
				    continue
             	}


        }
	}
}
