package greetings

import (
	"errors"
	"fmt"
)

//The Function hello needs a params of the name
func Hello(name string) (string, error) {
	// if no name
	if name == "" {
		return "", errors.New("No name given")
	}

	// Formats a string and returns it
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return "", message
}
