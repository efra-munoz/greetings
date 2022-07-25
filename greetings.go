package greetings

import (
	"errors"
	"fmt"
)

var ErrorEmptyName error = errors.New("empty name")

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf("Hi, %v. Welcome!", name), nil
}
