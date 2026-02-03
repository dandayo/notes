package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "time"
)
//Open file or create the new
func OpenFile(fileName string) *os.File{
	path := "notes/"+fileName+".txt"
	err := os.MkdirAll("notes", 0755)
	check(err)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()//Starting work with the file
	check(err)

	return file
}

//Read file for return list of notes
func ReadFile(fileName string) {
	path := "notes/"+fileName+".txt"
    myfile, err := os.Open(path)  //open the file in folder
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

//function to get the number of lines before append
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
//Add notes to file
func AddNote(fileName string, note string) {

	path := "notes/"+fileName+".txt" //check the path to folder with files
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE ,0644)//Add note in file, if we don't have file create a new one and check duplicate
	check(err)

    defer file.Close()

    size := fmt.Sprintf("%03d", fileLineCount(path) + 1)
    var data string

    timeStamp := time.Now().Format("02 Jan 2006 15:04")//Add timestamp for note

    if fileLineCount(path) != 0{
        data = "\n" + size + " - " + note + "	 " + timeStamp
    } else {
        data =  size + " - " + note + "		" + timeStamp //First note do not create a newline
    }
    if note == ""{

    } else {
    	_, err = file.WriteString(data)
    }
    check(err)
}

//Func for deleting file
func DeleteFile(fileName string) {
	path := "notes/"+fileName+".txt"
    err := os.Remove(path)
    check(err)
    fmt.Println("File deleted :(")
}
//Delete note by index from file
func RemoveNote(fileName string, removeIndex string) {
	path := "notes/"+fileName+".txt"
	intRemoveIndex, err := strconv.Atoi(removeIndex) //Convert input from string to int
	if err != nil {
		fmt.Println("\033[91mInvalid input, try again or use --help.\033[0m")
	}
	notes, err := os.Open(path)
	check(err)
    defer notes.Close()
    var lines []string
    scanner := bufio.NewScanner(notes)
    var index int = 0
    // for reordering of line number
    count := 0

    for scanner.Scan() {
         note := scanner.Text()
        //Skip file that we need to delete
         if index == intRemoveIndex - 1{
         	index++
          	continue
         }
         count++
         lineNumber := fmt.Sprintf("%03d", count)//Create a index tag for note
         lines = append(lines, lineNumber + " - " + note[6:])
         index++
    }
    if len(lines) < intRemoveIndex - 1{
    	fmt.Println(len(lines))
    	fmt.Println("\033[91mInvalid input, try again or use --help.\033[0m")
    }

    check(scanner.Err())
    update := strings.Join(lines, "\n")
    os.WriteFile(path, []byte(update), 0755)
}


    check(scanner.Err())
    update := strings.Join(lines, "\n")
    os.WriteFile(path, []byte(update), 0755)
}
