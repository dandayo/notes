package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CreateFile() {
    file, err := os.Create("new.txt")
    check(err)
    defer file.Close()
    fmt.Println("File created successfully")
}

func DeleteFile() {
    err := os.Remove("new.txt")
    check(err)
    fmt.Println("File deleted")
}

func WriteInFile() {
    file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
    check(err)
    defer file.Close()

    data := "new new info \n"
    _, err = file.WriteString(data)
    check(err)

}

func RemoveNote(fileName string, removeIndex int) {
	notes, err := os.Open(fileName)  //open the file
	check(err)
    defer notes.Close()

    var lines []string
    scanner := bufio.NewScanner(notes)
    //var index int = 0

    /*for scanner.Scan() {
    	if index != removeIndex {
     		lines = append(lines, scanner.Text())
     }
     index++
     } */



    for scanner.Scan() {
    	x := scanner.Text()
	    if (len(x)>= 0 || removeIndex-1 > len(x)){
			lines = append (x[:removeIndex], x[removeIndex+1:]...)
	    }
    }

    check(scanner.Err())

    update := strings.Join(lines, "\n")
    os.WriteFile(fileName, []byte(update), 0755)
}


func RemoveElement(array []string, n int)[]string{
		if (len (array)< 1|| n-1 > len (array)){
			return array
		}
		return append (array[:n], array[n+1:]...)
}

func main() {
 	RemoveNote("new.txt", 0)
}
