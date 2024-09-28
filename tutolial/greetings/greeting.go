package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// name is empty, return an error message
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(ranndomFormat(), name)
	return message, nil
}

// ranndomFormat returns on of set of greeting messages. the returnd message is selected at random.
func ranndomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
