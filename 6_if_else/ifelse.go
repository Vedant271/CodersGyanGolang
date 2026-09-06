package main

import "fmt"

func main() {
	// age := 20

	// if age >= 18 {
	// 	fmt.Println("Person is an adult", age)
	// } else {
	// 	fmt.Println("Person is not an adult", age)
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is an adult", age)
	// } else if age >= 11 {
	// 	fmt.Println("Person is a teenager", age)
	// } else {
	// 	fmt.Println("Person is a kid", age)
	// }

	role := "admin"
	hasPermissions := true

	if role == "admin" && hasPermissions {
		fmt.Println("Welcome aboard")
	}

	// if age := 2; age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is a teenager")
	// } else {
	// 	fmt.Println("Person is a kid")
	// }

	// Go does not contain any kind of ternary operator
}
