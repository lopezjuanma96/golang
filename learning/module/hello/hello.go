package main

import (
	"fmt"
	"log"

	/*
		for greetings.com/greetings to properly work, you need to run the following command:
		"go mod edit -replace greetings.com/greetings=../greetings"
		then the command "go mod tidy" to update the go.mod file, which will include
		a dependency on the local greetings module as "require greetings.com/greetings ..."
	*/
	"greetings.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, error := greetings.Hello("Juanma")
	// If an error was returned, print it to the console and exit the program.
	if error != nil {
		log.Fatal(error)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}
