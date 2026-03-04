package main

import (
	"fmt"
)

func Help() {
	fmt.Printf(`
Notes Tool - CLI Notes Manager

USAGE:
  ./notestool
  ./notestool --help

DESCRIPTION:
  After launch you can select an existing note file
  or create a new one using arrow keys.

NAVIGATION:
  ↑  Move up
  ↓  Move down
  Enter  Select option

WORKFLOW:
  1. Select existing file OR create new file
  2. Use main menu to manage notes

MAIN MENU OPTIONS:
  1. Show notes        - Display all notes in file
  2. Add a note        - Add new note to the end
  3. Delete a note     - Remove note by index
  4. Delete all notes  - Remove all notes (careful)
  5. Exit              - Close the program

NOTES:
  - All files are stored inside "notes/" directory
  - Files are saved in .txt format
  - Empty input is not allowed

EXAMPLE:
  ./notestool

`)
}
