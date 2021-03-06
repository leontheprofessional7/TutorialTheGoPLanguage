package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation) {
	// One cannot declare a variable but never use it
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "Hey!" + name
}

func main() {

	var salutation = Salutation{"Leon", "Hello!"}
	Greet(salutation)
}
