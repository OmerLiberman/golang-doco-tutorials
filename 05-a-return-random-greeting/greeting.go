package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Empty name")
	}

	message := fmt.Sprintf(randFormat(), name)
	return message, nil
}

func randFormat() string {
	formats := []string {
		"Hi, %v welcome!",
		"Great to see you !",
	}

	return formats[rand.Intn(len(formats))]
}

// Functions which are started with lower are accessible to modules withing the same package
