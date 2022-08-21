package main

import (
	"fmt"
	"log"
	"os"
)

var (
	BOLD   string = "\033[1m"
	YELLOW string = "\033[93m"
	BLUE   string = "\033[94m"
	END    string = "\033[0m"
)

func add(f int, s int) int {
	return f + s
}

func multi(f int, s int) int {
	return f * s
}

func main() {

	var firstNum int
	var secondNum int

	fmt.Printf("%sAncient simple Calculator%s\n\n", BOLD, END)

	fmt.Println("Enter the first digit: ")
	if _, err := fmt.Scan(&firstNum); err != nil {
		log.Print("Failed to scan first number", err)
		return
	}

	fmt.Println("Enter the second digit")
	if _, err := fmt.Scan(&secondNum); err != nil {
		log.Print("Failed to scan second number", err)
		return
	}

	fmt.Printf("%sSum     [ %d + %d = %d ]%s\n", BLUE, firstNum, secondNum, add(firstNum, secondNum), END)
	fmt.Printf("%sProduct [ %d * %d = %d ]%s\n", YELLOW, firstNum, secondNum, multi(firstNum, secondNum), END)
}
