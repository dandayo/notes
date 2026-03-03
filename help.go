package main

import (
	"fmt"
)

func Help() {
	fmt.Println(`
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

/*
Notes Tool — Help

This tool allows you to create, view, and manage your personal notes directly from the terminal.

How to start:
Run the program with a note file name:

    ./notestool note_name

If the file does not exist, it will be created automatically.
The provided name becomes the title of your notes file.

Navigation:
Use ↑ and ↓ arrow keys to move through the menu.
Press Enter to select an option.

Menu Options:

1. Show notes
Displays all saved notes for the selected file.
If there are no notes yet, you will see a message informing you.

2. Add a note
Allows you to write a new note.
After pressing Enter, the note is saved and you return to the main menu.

3. Delete a note
Shows the list of notes and asks you to enter the number of the note you want to remove.
Enter:
    • the note number to delete it
    • 0 to cancel the operation

4. Delete all notes (Be careful!)
Completely removes the current notes file.
You must confirm the action before deletion.

5. Exit
Closes the program safely.
All saved notes remain stored in the file.

Additional Notes:
• Invalid input will not crash the program.
• Empty input is not allowed.
• Notes are stored locally in your project directory.

Enjoy using Notes Tool!
*/
