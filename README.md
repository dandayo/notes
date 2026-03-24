# Notes tool

Secure CLI notes manager written in Go.

This tool allows you to create and manage note files directly from the terminal.
All notes are encrypted using the Blowfish algorithm and encoded with Base64 before being stored on disk.

## Usage
### How to Build
You need to build the tool before the start

```
$ go build -o notestool
```

Or you can change `notestool` to your own name

### How to run
You can start your note's tool in terminal. After you will see the selection where you can choose notes file or create new, it takes only first argument. If you want to use whitespace use "_" between words.

```
$ ./notestool
```
### Help

If you need to help or don't understand how it works write `--help` after.  
Also you can find the `help` in menu.

```
$ --help  or ./notestool --help
```

### Reset password

If you need to update your password, you can write  `--reset` with the `./notestool` (or your run word) and reset the password. Be careful you lost your previous notes, if you update your password.


```
$ ./notestool --reset
```

### Encryption

All notes are encrypted using the Blowfish algorithm and Base64 (to save it line by line) before being written to disk. 

* Encryption happens automatically when saving.  
* Decryption happens automatically when reading.  

Files stored in the `notes/` directory are not readable in plain text.  

Example of encrypted file content: 

```
iguO1xwYrvIMx7DMRDRoFD9EBy0zx498pAaIkukOHRTly3UFCja1rmeTefnqc4M2YFURqMgr8fe+hBBdsAT4vZAunQWz+cFTehRWERTv83/h70A=

```
Without the program, the content cannot be interpreted.

### Menu


1. Show notes.
> Show list of notes
2. Add a note.
> Add new note to the end of lists note
3. Delete a note.
> Delete the note by index, that you can see on the left side of note
4. Delete all notes.
> Delete the file with notes
5. Help.
> Show how to use program
6. Exit.
> To exit from tool

Use the arrows up, down and enter to see the result!

### Example
#### Welcome page
```
===Welcome to the notes tool!===

Select operation:

1. Show notes
2. Add a note
3. Delete a note
4. Delete all notes. Be careful!
5. Help
6. Exit
```
#### Add and show notes
```
Title: test
You don't have any notes! Add your first note

Select operation:

1. Show notes
-> Add a note
3. Delete a note
4. Delete all notes. Be careful!
5. Help
6. Exit

Enter the note text:
First note
```
```
Select operation:

-> Show notes.
2. Add a note.
3. Delete a note.
4. Delete all notes. Be careful
5. Help
6. Exit


Title: test
001 - First note		| 02 Feb 2026 11:12

```
#### Delete notes
```
Select operation:

1. Show notes.
2. Add a note.
-> Delete a note.
4. Delete all notes. Be careful
5. Help
6. Exit

Enter the number of note to remove or 0 to cancel:
1
```

#### Exit
```
===Goodbye!===
```


### Project Structure

```
notes/
├── notes/          # Folder with notes files
├── filesop.go      # File operations (create, read, delete notes)
├── menu.go         # Interactive menu interface
├── main.go         # Entry point of the application
├── secure.go       # Encrypt and decrypt functions
└── usage.go        # Often usage functions
```

### How Encryption Works

When you save a note:
1. The note text is encrypted using the Blowfish algorithm.
2. The encrypted binary data is encoded using Base64.
3. The Base64 string is written to the file.

When reading:
1. Base64 is decoded back into bytes.
2. Blowfish decryption is applied.
3. Original text is displayed.

### Mistakes

* Empty input.
* Wrong name on the the start
* Invalid input in menu

### Imports that I used in this tool

* bufio
* fmt
* os
* strconv
* strings
* time
* github.com/eiannone/keyboard (to use arrows)
* golang.org/x/crypto/blowfish
* encoding/base64
* crypto/cipher
* crypto/rand



## Author
__Daniil Danilov__



