package main

import (
	"fmt"
)

func Help() { //Print help for our user
	fmt.Print("\nNotes Tool — Help\n\nThis tool allows you to create, view, and manage your personal notes directly from the terminal.\n\nHow to start:\nRun the program with a note file name:\n  ./notestool note_name\nIf the file does not exist, it will be created automatically.\n\nNavigation:\nUse ↑ and ↓ arrow keys to move through the menu.\nPress Enter to select an option.\n\nMenu Options:\n1. Show notes — Displays all saved notes.\n2. Add a note — Write a new note and press Enter to save.\n3. Delete a note — Enter the note number to remove it or 0 to cancel.\n4. Delete all notes — Permanently removes the file after confirmation.\n5. Exit — Safely closes the program.\n\nInvalid input will not crash the program.\n")
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
