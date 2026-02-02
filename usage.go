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
		fmt.Println("\033[91mEmpty input.\nTry again!\033[0m")
	}
	if answer == "--help"{
		Help()
	}
	return answer
}

func EmptyTerminal() {
	fmt.Print("\033[3J\033[H\033[2J")
}

func check(e error) {

    if e != nil {
    	fmt.Println("\033[91mMistake! Something gone wrong\033[0m")
    }
}
