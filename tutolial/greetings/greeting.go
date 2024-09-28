package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// name is empty, return an error message
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
