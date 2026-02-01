package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
    	fmt.Println("Mistake!")
    }
}


func CreateFile(fileName string) {
     file, err := os.Create(fileName + ".txt")
     check(err)
     defer file.Close()
     fmt.Println("File created successfully")
}

func ReadFile(fileName string) {
    fmt.Println("\n===Current list===\n")

    myfile, err := os.Open(fileName + ".txt")  //open the file
    check(err)
    defer myfile.Close()

    scanner := bufio.NewScanner(myfile)  //scan the contents of a file and print line by line
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

    check(scanner.Err())
}


//function to get the number of lines before append
func fileLineCount(path string) int {
	file, err := os.Open(path)

	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

     lines := 0


	for scanner.Scan() {
	    lines++
            
	}

	return lines
}


func AddNote(fileName string, note string) {

	fileName = fileName + ".txt"
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY ,0755)

    defer file.Close()

    size := fmt.Sprintf("%03d", fileLineCount(fileName) + 1)
    var data string 


    

    if fileLineCount(fileName) != 0{
        data = "\n" + size + " - " + note
    } else {
        data =  size + " - " + note
    }

    
    _, err = file.WriteString(data)

    check(err)
}


func DeleteFile(fileName string) {
    err := os.Remove(fileName + ".txt")
    check(err)
    fmt.Println("File deleted")
}

func RemoveNote(fileName string, removeIndex string) {
	intRemoveIndex, err := strconv.Atoi(removeIndex)
	notes, err := os.Open(fileName + ".txt")  //open the file
	check(err)
    defer notes.Close()

    var lines []string
    scanner := bufio.NewScanner(notes)
    var index int = 0
    // for reordering of line number
    count := 0

    for scanner.Scan() {
         text := scanner.Text()
         
         if index == intRemoveIndex - 1{
         	index++
          	continue
         }

         count++
         lineNumber := fmt.Sprintf("%03d", count)
         lines = append(lines, lineNumber + " - " + text[6:])
         index++
    }



    check(scanner.Err())
    update := strings.Join(lines, "\n")
    os.WriteFile(fileName+".txt", []byte(update), 0755)


}

func input() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
   
	return strings.TrimSpace(scanner.Text())

}

func menu() {

	msg := ""
	name  := "new"

	for msg != "4" {

		fmt.Println("\n===Welcome to the notes tool!===\n")
		fmt.Println("Select operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		msg = input()

        switch msg {

            case "1":
                ReadFile(name)

            case "2":
                fmt.Println("\nEnter a note:")
                note := input()
			    AddNote(name, note)

            case "3":
                ReadFile(name)
                fmt.Println("\nEnter number to delete:")
                index := input()
			    RemoveNote(name, index)

            case "4":
                fmt.Println("===Goodbye!===\n")
			    os.Exit(0)

            default:
                fmt.Println("\n===You have entered an invalid input!===\n")
			    continue

        }

	}
}
func main() {
 	
    menu()
     
}
