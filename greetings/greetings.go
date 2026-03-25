package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Greetings returns a friendly greeting for the supplied name using a random template.
func Greetings(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

// Hello returns a deterministic greeting used in the tests.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hi, %v. Welcome!", name), nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
