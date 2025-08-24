package main

import "fmt"

const LoginToken = "djfioueiowndioiowesblWPYWBD" // This is public variable which must be first Letter Capital

func main() {
	var name string = "Aniket"
	fmt.Println(name)
	fmt.Printf("Type of variable is: %T \n", name)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable is: %T \n", isLoggedIn)

	var smallValue int8 = 123
	fmt.Println(smallValue)
	fmt.Printf("Type of variable is: %T \n", smallValue)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(LoginToken)
}
