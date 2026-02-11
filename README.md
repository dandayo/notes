# Notes tool

Here is our tool. You can create .txt files with list of notes, change note, show list of notes and delete note.

## Usage
### How to Build
 You need to build before the start
```
$ go build -o notestool
```

### How to run
You can start your note's in write terminal, it takes only first argument. If you want to use whitespace use "_" between words

```
$ ./notestool name_for_notes
```
You need to write name of your list of notes!
```
$ ./notestool  
Usage: ./todotool [TAG]  
```
### Help

If you need to help or don't understand how it works write `--help` after.

```
$ --help  
```

### Menu


1. Show notes.
> Show list of notes
2. Add a note.
> Add new note to the end of lists note
3. Delete a note.
> Delete the note by index, that you can see on the left side of note
4. Exit.
> To exit from tool

Just write the number of function and see the resut!

### Example
#### Welcome page
```
===Welcome to the notes tool!===

Select operation:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Delete all notes. Be careful
5. Exit.
```
#### Add and show notes
```
Title: test
You don't have any notes! Add your first note

Select operation:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Delete all notes. Be careful
5. Exit.

[YourInput] -> 2

Enter the note text:
First note

Select operation:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Delete all notes. Be careful
5. Exit.

[YourInput] -> 1

Title: test
001 - First note		02 Feb 2026 11:12

```
#### Delete notes
```
Select operation:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Delete all notes. Be careful
5. Exit.

[YourInput] -> 3
Enter the number of note to remove or 0 to cancel:
1

Select operation:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Delete all notes. Be careful
5. Exit.

[YourInput] -> 1

Title: test
You don't have any notes! Add your first note
```

#### Welcome page
```
[YourInput] -> 5
===Goodbye!===
```
### Mistakes

* Empty input.
* Wrong name on the the start
* Invalid input in menu

### Imports that we used in this tool

* bufio
* fmt
* os
* strconv
* strings
<<<<<<< HEAD
* time
=======
>>>>>>> ecb9897ad83d43b6eddcda62583a03e03fe9caf9



## Our team

<<<<<<< HEAD
__Daniil Danilov__    

=======
__Daniil Danilov__  
__Lauri Juhani Prusi__  
__Ryan Rodrigo Durante__  

__Thanks for using our tool!__
>>>>>>> ecb9897ad83d43b6eddcda62583a03e03fe9caf9
