package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ageCoparison(age1 int, age2 int) {
	if age1 > age2 {
		fmt.Println("Age1 is greater than Age2")
	} else if age1 < age2 {
		fmt.Println("Age1 is less than Age2")
	} else {
		fmt.Println("Age1 is equal to Age2")
	}
}
func switchCase(day string) {
	switch day {
	case "Monday":
		fmt.Printf("It is Monday")
	case "Tuesday":
		fmt.Printf("It is Tuesday")
	case "Wednesday":
		fmt.Printf("It is Wednesday")
	case "Thusday":
		fmt.Printf("It is Thusday")
	case "Friday":
		fmt.Printf("It is Friday")
	case "Saturday":
		fmt.Printf("It is Saturday")
	case "Sunday":
		fmt.Printf("It is Sunday")
	default:
		fmt.Printf("Invalid Input")
	}
}
func main() {
	var name string
	var age1, age2 int
	name = "Aniket"
	fmt.Print("Please enter your name")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
	fmt.Println("Enter two age in number")
	fmt.Scan(&age1, &age2)
	fmt.Scanln()
	ageCoparison(age1, age2)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the day: ")
	day, _ := reader.ReadString('\n')
	day = strings.TrimSpace(day)
	switchCase(day)

}
