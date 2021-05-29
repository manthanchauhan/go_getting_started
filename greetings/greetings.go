package greetings

import "fmt"

// only capital letter names can be exported

func Hello(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// or, message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
