package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Open file
func OpenFile(fileName string) *os.File {
	path := fileName
	if !strings.HasPrefix(path, "notes/") {
		path = "notes/" + path
	}
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}

	err := os.MkdirAll("notes", 0700)
	check(err)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0700)
	check(err)
	defer file.Close() //Starting work with the file

	return file
}

func CreateFile() string {
	fmt.Println("\n\033[32mCreate a new notes! If you want to go back input 0 \nChoose your name:\033[0m")
	cursorOn()
	newName := input()
	if newName == "" {
		CreateFile()
	} else if newName == "0" {
		main()
	}
	path := newName
	if !strings.HasPrefix(path, "notes/") {
		path = "notes/" + path
	}
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0700)
	check(err)
	defer file.Close()

	return newName
}

// Read file for return list of notes
func ReadFile(fileName string) {
	path := fileName
	if !strings.HasPrefix(path, "notes/") {
		path = "notes/" + path
	}
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}

	file, err := os.Open(path) //open the file in folder
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file) //scan the contents of a file and print line by line
	isEmpty := true
	for scanner.Scan() {
		isEmpty = false
		line := scanner.Text()
		fmt.Println(line)
	}
	if isEmpty { //If it empty return the message
		fmt.Println("\033[31mYou don't have any notes! Add your first note\033[0m")
	}
	check(scanner.Err())

}

// function to get the number of lines before append
func fileLineCount(path string) int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	return lines
}

const (
	TimeColor = "\033[38;5;105m"
	Reset     = "\033[0m"
)

func AddNote(fileName string, note string) {

	path := fileName
	if !strings.HasPrefix(path, "notes/") { // add this logic to avoid problems if user write empty input 2 times, dont create "notes/notes/fileName.txt.txt"
		path = "notes/" + path
	}
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}
	cursorOn()

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0700)
	check(err)
	defer file.Close()

	if note == "" || note == "--help" {
		return
	}

	lineNum := fileLineCount(path) + 1
	size := fmt.Sprintf("%03d", lineNum)

	timeStamp := time.Now().Format("02 Jan 2006 15:04")

	data := fmt.Sprintf(
		"%-3s - %-40s | %s%s%s",
		size,
		note,
		TimeColor, timeStamp, Reset,
	)

	// Move the string to next line if thats not first
	if lineNum > 1 {
		data = "\n" + data
	}

	_, err = file.WriteString(data)
	check(err)
}

// Func for deleting file
func DeleteFile(fileName string) {
	path := fileName
	if !strings.HasPrefix(path, "notes/") {
		path = "notes/" + path
	}
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}
	err := os.Remove(path)
	check(err)
	fmt.Println("\n\033[97;41mFile deleted :(\033[0m")
	fmt.Println("\033[42mGoodbye!\033[0m\n")
}

// Delete note by index from file
func RemoveNote(fileName string, removeIndex string) {
	path := fileName
	if !strings.HasPrefix(path, "notes/") {
		path = "notes/" + path
	}
	if !strings.HasSuffix(path, ".txt") {
		path += ".txt"
	}
	fileCount := fileLineCount(path)
	intRemoveIndex, err := strconv.Atoi(removeIndex) //Convert input from string to int
	if err != nil && removeIndex != "" {
		fmt.Println("\033[91mIndex should be a number!\033[0m")
	}
	notes, err := os.Open(path)
	check(err)
	defer notes.Close()

	var lines []string
	scanner := bufio.NewScanner(notes)
	var index int = 0
	// for reordering of line number
	count := 0

	if intRemoveIndex < 1 || intRemoveIndex > fileCount {
		fmt.Println("\033[91mIncorrect index! Try again.\033[0m")
	} else {
		for scanner.Scan() {
			note := scanner.Text()

			//Skip file that we need to delete
			if index == intRemoveIndex-1 {
				index++
				continue
			}
			count++
			lineNumber := fmt.Sprintf("%03d", count) //Create a index tag for note
			lines = append(lines, lineNumber+" - "+note[6:])
			index++

		}
		check(scanner.Err())
		update := strings.Join(lines, "\n")
		os.WriteFile(path, []byte(update), 0700)
	}
}
