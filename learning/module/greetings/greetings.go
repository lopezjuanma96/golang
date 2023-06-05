package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty_name")
	}
	// Embeds the name in a greeting message which is returned.
	message := fmt.Sprintf(randomFormat(), name)
	/*
		In Go, the := operator is a shortcut for declaring and initializing a variable in one line
		(Go uses the value on the right to determine the variable's type).
		Taking the long way, you might have written this as:
		var message string
		message = fmt.Sprintf(randomFormat(), name)
	*/

	return message, nil // A nil value indicates that there was no error.
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
