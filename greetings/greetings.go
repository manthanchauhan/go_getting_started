package greetings

import (
	"errors"
	"fmt"
)

// only capital letter names can be exported

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	// or, message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
