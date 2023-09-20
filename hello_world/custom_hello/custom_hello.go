package custom_hello

import (
	"fmt"
	"errors"
	"math/rand"
)

func CustomHello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string {
		"Hi, %v!, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := CustomHello(name)
		
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}
