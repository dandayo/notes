package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Open file or create the new
func OpenFile(fileName string) *os.File {
	path := "notes/" + fileName + ".txt"
	err := os.MkdirAll("notes", 0755)
	check(err)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close() //Starting work with the file
	check(err)

	return file
}

// Read file for return list of notes
func ReadFile(fileName string) {
	path := "notes/" + fileName + ".txt"
	myfile, err := os.Open(path) //open the file in folder
	check(err)
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile) //scan the contents of a file and print line by line
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

	path := "notes/" + fileName + ".txt"

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	defer file.Close()

	if note == "" {
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
	path := "notes/" + fileName + ".txt"
	err := os.Remove(path)
	check(err)
	fmt.Println("File deleted :(")
}

// Delete note by index from file
func RemoveNote(fileName string, removeIndex string) {
	path := "notes/" + fileName + ".txt"
	intRemoveIndex, err := strconv.Atoi(removeIndex) //Convert input from string to int
	notes, err := os.Open(path)
	check(err)
	defer notes.Close()

	var lines []string
	scanner := bufio.NewScanner(notes)
	var index int = 0
	// for reordering of line number
	count := 0
	fileCount := fileLineCount(path)

	if intRemoveIndex-1 < 0 || intRemoveIndex-1 > fileCount {
		fmt.Println("\033[31mIncorrect index! Try again.\033[0m")
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
		os.WriteFile(path, []byte(update), 0755)
	}
}
