package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
)

func input() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := strings.TrimSpace(scanner.Text())
	if answer == ""{
		fmt.Println("\033[31mEmpty input.\nTry again!\033[0m")
	}
	return answer
}

func EmptyTerminal() {
	fmt.Print("\033[3J\033[H\033[2J")
}

func check(e error) {

    if e != nil {
    	fmt.Println("\033[31mMistake!\033[0m")
    }
}
