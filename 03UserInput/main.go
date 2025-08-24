package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to online service")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for ranting, ", input)

}
