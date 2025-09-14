package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func switchCase(day string) {
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Almost weekend")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek day")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the day: ")
	name, _ := reader.ReadString('\n')

	// Trim newline and extra spaces
	name = strings.TrimSpace(name)

	switchCase(name)
}
