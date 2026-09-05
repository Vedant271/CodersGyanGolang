package main

import "fmt"

const newName = "new golang"
const anotherNewName string = "another new golang"

var age = 20
var newAge int = 40

const (
	port = 5000
	host = "localhost"
)

const (
	databaseUsername string = "user"
	databasePassword int    = 1234
)

func main() {
	const name = "golang"

	fmt.Println(name)
	fmt.Println(newName)
	fmt.Println(age)
	fmt.Println(newAge)

	fmt.Println(port)
	fmt.Println(host)

	fmt.Println(databaseUsername, databasePassword)
}
