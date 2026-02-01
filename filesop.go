package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "time"
)

func OpenFile(fileName string) *os.File{
	path := "notes/"+fileName+".txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	return file
}

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
    if isEmpty {
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

func AddNote(fileName string, note string) {

	path := "notes/"+fileName+".txt"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE ,0644)
	check(err)

    defer file.Close()

    size := fmt.Sprintf("%03d", fileLineCount(path) + 1)
    var data string

    timeStamp := time.Now().Format("02 Jan 2006 15:04")

    if fileLineCount(path) != 0{
        data = "\n" + size + " - " + note + "	 " + timeStamp
    } else {
        data =  size + " - " + note + "		" + timeStamp
    }
    if note == ""{
        fmt.Println("\033[31mEmpty input.\033[0m")

    } else {
    	_, err = file.WriteString(data)
    }
    check(err)
}


func DeleteFile(fileName string) {
    err := os.Remove(fileName)
    check(err)
    fmt.Println("File deleted")
}

func RemoveNote(fileName string, removeIndex string) {
	path := "notes/"+fileName+".txt"
	intRemoveIndex, err := strconv.Atoi(removeIndex)
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

         if index == intRemoveIndex - 1{
         	index++
          	continue
         }

         count++
         lineNumber := fmt.Sprintf("%03d", count)
         lines = append(lines, lineNumber + " - " + note[6:])
         index++
    }

    check(scanner.Err())
    update := strings.Join(lines, "\n")
    os.WriteFile(path, []byte(update), 0755)
}
