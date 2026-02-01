package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
	fmt.Println("\n===Welcome to the notes tool!===\n")
	fileName := os.Args[1]
	OpenFile(fileName)
	menu(fileName)
}

func OpenFile(name string) *os.File{
	path := "notes/"+name+".txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	return file
}

func ReadFile(name string) {
	path := "notes/"+name+".txt"
    myfile, err := os.Open(path)  //open the file in folder
    check(err)
    defer myfile.Close()

    scanner := bufio.NewScanner(myfile)
    notes := scanner.Scan()  //scan the contents of a file and print line by line
    for notes {
        line := scanner.Text()
        fmt.Println(line)
    }

    check(scanner.Err())
}

func check(e error) {

    if e != nil {
    	fmt.Println("Mistake!")
    }
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

    if fileLineCount(path) != 0{
        data = "\n" + size + " - " + note
    } else {
        data =  size + " - " + note
    }

    if note != ""{
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
    os.WriteFile(path, []byte(update), 0644)
}

func input() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := strings.TrimSpace(scanner.Text())
	if answer == ""{
		fmt.Println("Empty input.\nTry again!")
	}

	return answer
}

func menu(name string) {

	msg := ""

	for msg != "4" {
		fmt.Println("\n")
		fmt.Println("Select operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")
		fmt.Println("\n")

		msg = input()

        switch msg {

            case "1":
            	fmt.Println("===Your notes!===\n")
                ReadFile(name)

            case "2":
            	fmt.Println("Enter the note text:")
                note := input()
			    AddNote(name, note)

            case "3":
            	fmt.Println("Enter the number of note to remove or 0 to cancel:")
                index := input()
                if index == "0" {
                continue
                } else {
                	RemoveNote(name, index)
                }

            case "4":
                fmt.Println("===Goodbye!===\n")
			    os.Exit(0)

            default:
            	if msg != ""{
	             	fmt.Println("\n===You have entered an invalid input!===\n")
				    continue
             	}


        }

	}
}
